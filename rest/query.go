package rest

//
//import (
//    "github.com/gin-gonic/gin"
//    "github.com/irisnet/iris-community/models"
//    "net/http"
//    "fmt"
//)
//
//type Result struct{
//    Id         uint     //用户ID
//    Code       string
//    Invite     uint    //总邀请人
//    Complete   uint    //验证通过的邀请人
//    Sum        int      //获得token 总数
//}
//
//func QueryRegister(g *gin.RouterGroup) {
//    g.GET("", QueryInfo)
//}
//
//
//func QueryInfo(c *gin.Context) {
//    authCode := c.Request.Header.Get("Authorization")
//
//    userAuth := models.UserAuth{
//        AuthCode: authCode,
//    }
//
//    userAuth.FindByAuth()
//
//    id := userAuth.UserId
//
//    //id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
//       //get invitation_code
//      var target models.UserProfile
//       models.DB.Where("user_id = ?", id).First(&target)
//
//       //fmt.Println("code  : ", target.InvitationCode)
//
//        //query join result
//        result := Result{Id:id,Code:target.InvitationCode}
//        rows, err := models.DB.Table("user_invitation").Select("user_approval.user_id,user_invitation.invitation_code, user_approval.approval_status").Joins("left join user_approval on user_approval.user_id = user_invitation.invitee_id").Where("user_invitation.invitation_code = ?",target.InvitationCode).Rows()
//
//        if err != nil{
//            panic(err)
//        }
//
//        //获得token总数
//        sum := 0
//        for rows.Next() {
//
//            var id int //用户ID
//            var invitation_code string //用户邀请码
//            var is_approved string
//            rows.Scan(&id,&invitation_code,&is_approved)
//            fmt.Println(id,invitation_code,is_approved)
//            result.Invite++
//
//            if is_approved == "p" {
//
//                //完成注册人数增加
//                result.Complete ++
//                //获得token 数量
//                rows, err := models.DB.Table("user_x_token").Select("dim_token_access_mode.num").Joins("left join dim_token_access_mode on dim_token_access_mode.id = user_x_token.access_mode_id").Where("user_x_token.user_id = ?",id).Rows()
//                if err != nil{
//                    print(err)
//                }
//                var num int
//                for rows.Next() {
//                    rows.Scan(&num)
//                    sum += num
//                    //fmt.Println(num)
//                }
//            }
//
//        }
//
//        result.Sum = sum
//
//        c.JSON(http.StatusOK, result)
//
//}
