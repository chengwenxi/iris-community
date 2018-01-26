package rest

import (
    "github.com/gin-gonic/gin"
    "github.com/irisnet/iris-community/models"
    "strconv"
    "net/http"
)

type Result struct{
    Id         uint64     //用户ID
    Code       string
    Invite     uint    //验证通过的邀请人
    Complete        uint    //总邀请人

}

func QueryRegister(g *gin.RouterGroup) {
    g.GET("/:id", QueryInfo)
}


func QueryInfo(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
       //get invitation_code
      var target models.UserProfile
       models.DB.Where("user_id = ?", id).First(&target)

       //fmt.Println("code  : ", target.InvitationCode)
        
        //query join result
        rows, err := models.DB.Table("user_invitation").Select("user_invitation.invitation_code, user_approval.approval_status").Joins("left join user_approval on user_approval.user_id = user_invitation.invitee_id").Where("user_invitation.invitation_code = ?",target.InvitationCode).Rows()
   
        if err != nil{
            print(err)
        }

        result := Result{Id:id,Code:target.InvitationCode}

        for rows.Next() {
           
            var invitation_code string
            var is_approved string
            rows.Scan(&invitation_code,&is_approved)
            //fmt.Println(invitation_code,is_actived)

            result.Invite++

            if is_approved == "f" {
                result.Complete ++
            }
        }

       
        c.JSON(http.StatusOK, result)

}
