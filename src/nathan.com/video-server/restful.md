## user module
#### 创建（注册）用户
- url:/user
- method:post
- status code:201 400 500

#### 用户登录
- url:/user/:username
- method:post
- status code:200 400 500

#### 获取用户基本信息
- url:/user/:username
- method:get
- status code:200 400 401 403 500

#### 获取用户基本信息
- url:/user/:username
- method:delete
- status code:204 400 401 403 500

## 用户资源
#### list all video_ideos
- url:/user/:username/video_ideos
- method: get
- status code: 200 400 500

#### get one video_ideos
- url:/user/:username/video_ideos/:video_id
- method: get
- status code: 200 400 500

#### delete one video_ideos
- url:/user/:username/video_ideos/:video_id
- method: delete
- status code: 204 400 401 402 500

## 评论
#### show comments
- url: /video_ideos/:video_id/comments
- method: get
- status code: 200 400 500

#### post a comment
- url: /video_ideos/:video_id/comments
- method: post
- status code: 201 400 500

#### delete a comment
- url: /video_ideos/:video_id/comments/:comment_id
- method: delete
- status code: 204 400 401 402 500
