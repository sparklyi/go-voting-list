下载后配置config文件夹下的mysql.go文件和redis文件  
本项目涉及到的技术栈:**gin+gorm+logrus+go-redis+MySQL+Vue3**  
采用宝塔面板搭载**nginx+redis+MySQL**运行  

利用loggrus自定义日志，实现不同等级日志记录及访问记录，快速恢复
利用redis优化排行榜响应速度

After downloading, configure the mysql.go file and redis file in the config folder  

Technical stack involved in this project :**gin+gorm+logrus+go-redis+MySQL+Vue3**  
Using the pagoda panel with **nginx+redis+MySQL** running  

Use loggrus custom logs to achieve different levels of logging and access records for fast recovery  
Optimize leaderboard responsiveness with redis  

项目结构
````
├─cache                  //缓存文件夹
│      redis.go          //启动redis等操作
│      
├─config                         //配置信息文件夹
│      mysql.go			 //MySQL的配置信息，passwd，port等
│      redis.go			 //redis的配置信息，address，db等
│      
├─controller			 //控制器文件夹
│      activity.go		 //处理活动相关请求的函数文件
│      common.go		 //公用的一些函数，如MD5加密等
│      player.go		 //处理候选人相关请求的函数文件
│      user.go			 //处理用户相关请求的函数，如login、register等
│      vote.go			 //处理投票相关请求的函数，如addVote，getRanking等
│      
├─dao			         //数据访问对象
│      dao.go			 //完成数据库的连接，连接池的设置等
│      
├─models		         //模型（表）
│      activity.go		 //活动模型，模型结构，相关操作
│      player.go		 //选手模型
│      user.go			 //用户模型
│      vote.go			 //投票模型
│      
├─pkg		                 //包
│  └─logger				
│          logger.go	         //自定义日志记录，支持不同等级的错误，以及recover处理
│          
├─router                         //路由
│      routers.go		 //处理发起的各类请求，连接controller处理
│      
└─runtime                        //运行
    └─log                        //存放记录的日志信息
````  
参考视频：[下雨le](https://www.imooc.com/learn/)
