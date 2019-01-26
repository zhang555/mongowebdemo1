//
db.user.aggregate(
    [
        {
            $group: {
                _id: "$roleName",
                roleNameShow: {$push: "$roleNameShow"},
                count: {$sum: 1},
                usernames: {$push: "$username"},
            }
        }
    ]
).pretty()


db.user.aggregate(
    [
        {
            $group: {
                _id: {
                    roleName: "$roleName",
                    roleNameShow: "$roleNameShow",
                },
                count: {$sum: 1},
                usernames: {$push: "$username"},
            }
        }
    ]
).pretty()


db.user.update(
    {'roleName': 'user'},
    {$set: {'roleNameShow': '用户'}},
    {multi: true}
)

db.user.update(
    {'roleName': 'admin'},
    {$set: {'roleNameShow': 'admin'}},
    {multi: true}
)


db.user.find({}, {
    "roleName": 1,
    "roleNameShow": 1,
}).pretty()



//报错，那个逗号不能加，
//有的地方多个逗号没事，这个地方多逗号报错，
db.user.find(
    {"roleName": "user",},
    {"roleNameShow": 1,},
).pretty()


db.user.find(
    {"roleName": "user",},
    {"roleNameShow": 1,}
).pretty()



