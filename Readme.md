## /CreateChat - (FirstUserId, SecondUserId) -> code 200

## /GetChats - (Id (userId)) -> [{Id, FirstUserId, SecondUserId}, ...]

## /Login - (Secret) -> {Id, Secret, Image, Name, Address}

## /SendMessage - (Text, UserId, ChatId) -> code 200

## /GetMessagesInChat - (Id (ChatId)) -> [{Id, Text, UserId, ChatId}, ...]

## /GetUser - (Id (UserId)) -> {Id, Secret, Name, Image, Address}
