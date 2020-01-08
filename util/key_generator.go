package util

import "strconv"

// generate a key for redis  for specific app chat to save this chat's info (last msg number, this chat's id ..etc)
// cause there could be many chats with same number but belongs to different apps
func GenerateAppChatKey(appToken string, chatNum uint) string {
	chatNumStr := strconv.Itoa(int(chatNum))
	appChatKey := appToken + "-" + chatNumStr
	return appChatKey
}
