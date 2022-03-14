package types

import "fmt"

// short identifiers for audit actions
const (
	ActionShareCreated            = "file_shared"
	ActionSharePermissionUpdated  = "share_permission_updated"
	ActionShareDisplayNameUpdated = "share_name_updated"
	ActionSharePasswordUpdated    = "share_password_updated"
	ActionShareExpirationUpdated  = "share_expiration_updated"
	ActionShareRemoved            = "file_unshared"
	ActionShareAccepted           = "share_accepted"
	ActionShareDeclined           = "share_declined"
	ActionLinkAccessed            = "public_link_accessed"
)

// MessageShareCreated returns the human readable string that describes the action
func MessageShareCreated(sharer, item, grantee string) string {
	return fmt.Sprintf("user '%s' shared file '%s' with '%s'", sharer, item, grantee)
}

// MessageLinkCreated returns the human readable string that describes the action
func MessageLinkCreated(sharer, item, shareid string) string {
	return fmt.Sprintf("user '%s' created a public to file '%s' with id '%s'", sharer, item, shareid)
}

// MessageShareUpdated returns the human readable string that describes the action
func MessageShareUpdated(sharer, shareID, fieldUpdated string) string {
	return fmt.Sprintf("user '%s' updated field '%s' of share '%s'", sharer, fieldUpdated, shareID)
}

// MessageLinkUpdated returns the human readable string that describes the action
func MessageLinkUpdated(sharer, shareid, fieldUpdated string) string {
	return fmt.Sprintf("user '%s' updated field '%s' of public link '%s'", sharer, fieldUpdated, shareid)
}

// MessageShareRemoved returns the human readable string that describes the action
func MessageShareRemoved(sharer, shareid, itemid string) string {
	return fmt.Sprintf("share id:'%s' uid:'%s' item-id:'%s' was removed", shareid, sharer, itemid)
}

// MessageLinkRemoved returns the human readable string that describes the action
func MessageLinkRemoved(shareid string) string {
	return fmt.Sprintf("public link id:'%s' was removed", shareid)
}

// MessageShareAccepted returns the human readable string that describes the action
func MessageShareAccepted(userid, shareid, sharerid string) string {
	return fmt.Sprintf("user '%s' accepted share '%s' from user '%s'", userid, shareid, sharerid)
}

// MessageShareDeclined returns the human readable string that describes the action
func MessageShareDeclined(userid, shareid, sharerid string) string {
	return fmt.Sprintf("user '%s' declined share '%s' from user '%s'", userid, shareid, sharerid)
}

// MessageLinkAccessed returns the human readable string that describes the action
func MessageLinkAccessed(linkid string, success bool) string {
	return fmt.Sprintf("link '%s' was accessed. Success: %v", linkid, success)
}