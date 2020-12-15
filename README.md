# duck
亂數產生package

使用linux提供的系統產生亂數方式，srand() 產生亂數，因此windows系統不適用此package。

srand 產生需要一點時間(看硬體效能)，因此預設有10個goroutine輔助併發跟系統取亂數後，塞進channel儲備。

使用時，是從儲備的亂數channel取得。


## 目前現有亂數庫存數量
GetRandStorageNum

## 取得mod 條件的亂數
GetRandUnderRange

##  取得指定 min ~ max的亂數
GetRandBetweenRange