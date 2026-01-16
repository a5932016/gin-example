# Database
新增的Table加到DB.AutoMigrate()裡面後可以自動創建
如果需要初始化、測試資料可以寫到seed function裡面

## ORM
- [connecting to database](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)

### 優點
1) 簡化代碼：ORM將數據庫操作映射到對象操作，可以更直觀地進行數據庫操作，從而減少了編寫SQL語句的需要，簡化了代碼。
2) 對象導向：ORM支持對象導向編程，使得數據庫操作更符合面向對象的開發模型，對於開發人員來說更容易理解和維護。
3) 跨平台兼容性：ORM通常是跨平台的，可以在不同的數據庫系統中使用，而不需要修改底層SQL語句。

### 缺點
1) 學習曲線：學習和掌握ORM框架需要一定的時間和學習成本，尤其是對於複雜的映射關係和高級功能，可能需要更多的學習和實踐。
2) 效能限制：由於ORM需要對象映射和轉換，可能在一些性能要求較高的場景下出現性能瓶頸。有時候直接使用SQL可以更有效地編寫和優化特定的查詢語句。
3) 映射限制：ORM在對象和數據庫之間的映射中可能存在一些限制，特別是當存在複雜的數據模型和關係時。某些特定的數據庫功能可能無法完全由ORM支持。

## SQL language
- [postgresql](https://docs.postgresql.tw/the-sql-language)

## 優點
1) 強大的查詢能力：SQL是用於數據庫操作的標準語言，具有豐富的查詢功能和靈活性。可以進行複雜的數據查詢和聚合操作。
2) 高效性能：直接使用SQL可以對數據庫進行細粒度的控制，優化查詢語句，從而獲得更高的性能和效率。
3) 完整的數據庫功能支持：SQL可以直接調用數據庫提供的各種功能和特性，例如事務處理、索引、觸發器等。

## 缺點
1) 學習曲線：相對於ORM，學習和掌握SQL的語法和操作需要更多的時間和努力。對於不熟悉SQL的開發人員來說，創建和優化複雜的查詢可能會具有一定的難度。
2) 跨平台兼容性：SQL是一種標準語言，但不同的數據庫系統可能具有不同的SQL方言和特性。這可能導致在不同的數據庫之間遷移和運行SQL語句時出現一些兼容性問題。
3) 硬編碼：使用SQL時，需要直接在代碼中硬編碼SQL語句。這導致了代碼的耦合性增加，並且在需要更改查詢時可能需要修改代碼，增加了維護的困難性。
4) 安全性風險：直接使用SQL語句可能存在安全風險，例如SQL注入攻擊。如果沒有適當地處理用戶輸入，攻擊者可能通過構造惡意的SQL語句來獲取或修改數據庫中的數據。
5) 數據庫依賴性：直接使用SQL可能導致代碼與特定的數據庫系統綁定。如果需要更換數據庫系統，可能需要重寫和調整現有的SQL語句，增加了代碼的耦合性和遷移的困難性。

## database skill
- [transaction isolation](https://docs.postgresql.tw/v/13/the-sql-language/concurrency-control/transaction-isolation)
- [database normalization](https://zh.wikipedia.org/zh-tw/%E6%95%B0%E6%8D%AE%E5%BA%93%E8%A7%84%E8%8C%83%E5%8C%96)
- [seed](https://en.wikipedia.org/wiki/Database_seeding)
- [migration](https://gorm.io/zh_CN/docs/migration.html)

## databases
- [sqlite](https://www.sqlite.org/index.html)
- [postgresql](https://docs.postgresql.tw/)

## database studio
- [sqlite](https://sqlitestudio.pl/)
- [postgresql](https://www.postgresqlstudio.org/)