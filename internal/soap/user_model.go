package soap

import "encoding/xml"

// UserResponse UserInfo represents the root XML structure for the user info response
type UserResponse struct {
	XMLName    xml.Name   `xml:"userInfo"`
	UserDetail UserDetail `xml:"userDetail"`
}

// UserDetail holds the details of the user
type UserDetail struct {
	UserID      string `xml:"userID"`
	UserName    string `xml:"userName"`
	Email       string `xml:"email"`
	PhoneNumber string `xml:"phoneNumber"`
}
