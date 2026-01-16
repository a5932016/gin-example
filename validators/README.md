# Validator
此資料夾為自訂義Validator
定義好的規則可以在request model 的 binding:"規則" 便可使用
有一些預設的規則可以在[Validator V10](https://github.com/go-playground/validator)裡面找
要注意如果有設定binding規則，該欄位會自動設為必填，可以使用| ex. binding:"email|empty" 意思就是允許空或Email

[Gin Validator](https://gin-gonic.com/docs/examples/custom-validators/)