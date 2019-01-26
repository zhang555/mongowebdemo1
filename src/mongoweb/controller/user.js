db.user.find({
        _id: ObjectId("5c4027a59816112290e4f329"),
    },
    {
        username: 1,
        roleName: 1,
        roleNameShow: 1,
        createTime: 1,
        createTimeInt: 1,
        updateTime: 1,
        updateTimeInt: 1,

    }
).pretty()

db.user.find({
        _id: ObjectId("5c4027a59816112290e4f329"),
    }
).pretty()


