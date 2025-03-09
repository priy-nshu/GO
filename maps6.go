package main

import(
  "slices"
)

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggestions := make(map[string]bool) 
	userFriends, exists := friendships[username]

	if !exists {
		return nil 
	}
	
	for _, friend := range userFriends {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username && !slices.Contains(userFriends, friendOfFriend) {
				suggestions[friendOfFriend] = true
			}
		}
	}
	if len(suggestions) == 0 {
		return nil 
	}

	result := make([]string, 0, len(suggestions))
	for friend := range suggestions {
		result = append(result, friend)
	}

	return result
}
