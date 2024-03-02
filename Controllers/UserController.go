package Controllers


func Signup(c *gin.Context) {
	var user Models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while hashing password!"})
		return
	}
	user.Password = string(hash)

	if err := Database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not created!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sucess": "User created successfully!"})
}