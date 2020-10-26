#! /bin/bash

# build web ui
cd /Users/bytedance/stream-player/src/nathan.com/video-server/web
go install
cp -r /Users/bytedance/stream-player/bin/web /Users/bytedance/stream-player/video_server_web_ui/bin/web
cp -r /Users/bytedance/stream-player/src/nathan.com/video-server/templates /Users/bytedance/stream-player/video_server_web_ui/bin/web/templates