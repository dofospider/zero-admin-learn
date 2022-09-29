# zero-admin-learn
### 起因：
  因爲从零开始学习go的微服务，不瞭解整體的架構以及各個不見的工作機理。於是想着找一個前後端都有的項目跑一下，這樣就可以知道整體的工作機理了。想法很簡單過程很曲折，從開始嘗試搭建到今天正式完成將近用了3個禮拜的業餘時間。

### 使用步驟：
  #### 1.使用zero-admin中的程序構建docker鏡像。
            cd zero-admin
            docker build -t sys:v1 -f rpc/sys/Dockerfile .
            docker build -t ums:v1 -f rpc/ums/Dockerfile .
            docker build -t oms:v1 -f rpc/oms/Dockerfile .
            docker build -t pms:v1 -f rpc/pms/Dockerfile .
            docker build -t sms:v1 -f rpc/sms/Dockerfile .
            docker build -t api:v1 -f api/Dockerfile .

  #### 2. 使用zero-admin-docker中的yml文件構建dockers並給數據庫寫入數據。
            docker compose -f "docker-compose.yml" up -d --build
            
            使用數據庫管理程序使用root/123456連接mysql數據庫建立數據表空間gozero，
            導入gozero.sql。

  #### 3.進入zero-vue-admin
            npm install
            npm  run dev
            

  #### ps
            node.js v16.17.0
