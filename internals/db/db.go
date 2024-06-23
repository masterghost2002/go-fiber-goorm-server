package db

var storageUser = make(map[int]string)
var storagePost = make(map[int]map[int]string)

func GetUser(userId int) string {
	return storageUser[userId]
}
func GetPost(userId, postId int) string {
	return storagePost[userId][postId]
}
func AddUser(userId int, name string) {
	storageUser[userId] = name
}
