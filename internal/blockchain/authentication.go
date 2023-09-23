package blockchain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	caMsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// User stuct that allow a registered user to query and invoke the blockchain
type FabricUser struct {
	Username        string
	Fabric          *FabricSetup
	SigningIdentity msp.SigningIdentity
	ChannelClient   *channel.Client
}

var sessionUser = make(map[string]string)
var secretKey = make(map[string]string)

func (s *FabricSetup) SessionUser() (*FabricUser, error) {

	if len(sessionUser["name"]) > 0 {

		var err error

		var username = sessionUser["name"]

		var user FabricUser
		user.Username = username
		user.Fabric = s

		user.SigningIdentity, err = s.CaClient.GetSigningIdentity(username)
		if err != nil {
			return nil, fmt.Errorf("failed to get signing identity for '%s': %v", username, err)
		}

		clientContext := s.sdk.ChannelContext(
			s.ChannelID, fabsdk.WithUser(username),
			fabsdk.WithOrg(s.OrgName),
			fabsdk.WithIdentity(user.SigningIdentity))

		// Channel client is used to query and execute transactions
		user.ChannelClient, err = channel.New(clientContext)
		if err != nil {
			return nil, fmt.Errorf("failed to create new channel client for '%s': %v", username, err)
		}
		fmt.Println("Channel client created")

		// Creation of the client which will enables access to our channel events
		s.event, err = event.New(clientContext)
		if err != nil {
			return nil, fmt.Errorf("failed to create new event client %v", err)
		}

		return &user, nil
	}

	return nil, nil
}

func (s *FabricUser) Logout() {

	delete(sessionUser, "name")
}

func (s *FabricUser) RevokeUser(email string) error {

	_, err := s.Fabric.CaClient.Revoke(&caMsp.RevocationRequest{
		Name: email,
	})

	if err != nil {
		return fmt.Errorf("failed to revoke signing identity for '%s': %v", email, err)
	}

	return nil
}

func (s *FabricUser) RemoveUserFromCA(email string) error {

	_, err := s.Fabric.CaClient.RemoveIdentity(&caMsp.RemoveIdentityRequest{

		ID:     email,
		Force:  true,
		CAName: s.Fabric.CaID,
	})

	if err != nil {
		return fmt.Errorf("failed to remove signing identity for '%s': %v", email, err)
	}

	return nil
}

func (s *FabricUser) RemoveUser(email string) error {

	err := s.RemoveUserFromCA(email)

	if err != nil {
		return fmt.Errorf("failed to remove signing identity for '%s': %v", email, err)
	} else {

		err = s.DeleteUserFromLedger(email)
		if err != nil {
			return fmt.Errorf("unable to delete user from blockchain:>>> %v", err)
		}

	}

	return nil

}

func (u *FabricUser) UpdateUserData(userId, name, email, company, occupation, salary, userType string) error {

	_, err := u.UpdateUserDataFromLedger(userId, name, email, company, occupation, salary, userType)

	if err != nil {
		return fmt.Errorf("unable to update user data in blockchain:>>> %v", err)
	}

	return nil
}

func (s *FabricUser) ChangePassword(email, userType, oldPwd, newPwd string) error {

	if !strings.EqualFold(oldPwd, secretKey["secret"]) {
		return errors.New("failed old password not matched, can't change pwd")
	}

	if strings.EqualFold(oldPwd, newPwd) {
		return errors.New("failed old password, new password should not be same")
	}

	err := s.RemoveUserFromCA(email)

	if err != nil {
		return fmt.Errorf("failed to remove identity '%s': %v", email, err)
	}

	fmt.Println("User Removed")

	err = s.Fabric.RegisterWithCA(email, newPwd, userType)

	if err != nil {
		return fmt.Errorf("failed to register with CA '%s': %v", email, err)
	}

	fmt.Println("User Re-Registered")

	err = s.Fabric.EnrollUser(email, newPwd)

	if err != nil {
		return fmt.Errorf("failed to enroll user '%s': %v", email, err)
	}

	fmt.Println("User Enrolled")

	err = s.Fabric.ReEnrollUser(email)

	if err != nil {
		fmt.Println("Error Re-Enroll User : %s ", err.Error())
		return fmt.Errorf("failed to re-enroll user '%s': %v", email, err)
	}

	fmt.Println("User Re-Enrolled")

	return nil
}

func (s *FabricSetup) ReEnrollUser(email string) error {

	err := s.CaClient.Reenroll(email)

	if err != nil {
		return fmt.Errorf("failed to re-enroll user '%s': %v", email, err)
	}
	return nil
}

func (s *FabricSetup) EnrollUser(email, password string) error {

	err := s.CaClient.Enroll(email, caMsp.WithSecret(password))
	if err != nil {
		return fmt.Errorf("failed to enroll identity '%s': %v", email, err)
	}

	secretKey["secret"] = password

	return nil
}

func (s *FabricSetup) LoginUser(email, password string) (*FabricUser, error) {

	err := s.EnrollUser(email, password)

	if err != nil {
		return nil, fmt.Errorf("failed to enroll user '%s': %v", email, err)
	}

	var user FabricUser
	user.Username = email
	user.Fabric = s

	user.SigningIdentity, err = s.CaClient.GetSigningIdentity(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get signing identity for '%s': %v", email, err)
	}

	clientContext := s.sdk.ChannelContext(
		s.ChannelID, fabsdk.WithUser(email),
		fabsdk.WithOrg(s.OrgName),
		fabsdk.WithIdentity(user.SigningIdentity))

	// Channel client is used to query and execute transactions
	user.ChannelClient, err = channel.New(clientContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create new channel client for '%s': %v", email, err)
	}
	fmt.Println("Channel client created")

	// Creation of the client which will enables access to our channel events
	s.event, err = event.New(clientContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create new event client %v", err)
	}
	fmt.Println("Event client created")

	fmt.Println("Login Successful")

	sessionUser["name"] = email

	return &user, nil
}

func (s *FabricSetup) RegisterWithCA(email, password, userType string) error {

	fmt.Println("CA Register === " + s.CaID)

	_, err := s.CaClient.Register(&caMsp.RegistrationRequest{
		Name:           email,
		Secret:         password,
		Type:           "user",
		MaxEnrollments: -1,
		Affiliation:    "org1",
		Attributes: []caMsp.Attribute{
			{
				Name:  "usermode",
				Value: userType,
				ECert: true,
			},
		},
		CAName: s.CaID,
	})

	if err != nil {
		return fmt.Errorf("unable to register user with CA '%s': %v", email, err)
	}

	return nil
}

func (s *FabricSetup) RegisterUser(name, email, company, occupation, salary, password, userType string) (*FabricUser, error) {

	fmt.Println("****** Register User ****** " + name + " , " + email + " , " + company + " , " + occupation + " , " + salary + " , " + password + " , " + userType)

	err := s.RegisterWithCA(email, password, userType)

	if err != nil {
		return nil, fmt.Errorf("unable to register user '%s': %v", email, err)
	}

	u, err := s.LoginUser(email, password)
	if err != nil {
		return nil, fmt.Errorf("unable to log user '%s' after registration: %v", email, err)
	}

	err = u.addUserToLedger(name, email, company, occupation, salary, userType)

	if err != nil {
		return nil, fmt.Errorf("unable to add user to blockchain:>>> %v", err)
	}

	return u, nil
}
