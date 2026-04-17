package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CertificateController struct {
	CertificateService services.CertificateService
}

func NewCertificateController(certificateService services.CertificateService) *CertificateController {
	return &CertificateController{CertificateService: certificateService}
}

func (cc *CertificateController) CreateCertificate(c *gin.Context) {
	var certificate models.Certificate
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.CertificateService.CreateCertificate(&certificate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificate created successfully"})
}

func (cc *CertificateController) UpdateCertificate(c *gin.Context) {
	var certificate models.Certificate
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.CertificateService.UpdateCertificate(&certificate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificate updated successfully"})
}

func (cc *CertificateController) DeleteCertificate(c *gin.Context) {
	var certificate models.Certificate
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.CertificateService.DeleteCertificate(&certificate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificate deleted successfully"})
}

func (cc *CertificateController) GetCertificateByID(c *gin.Context) {
	var certificate models.Certificate
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := cc.CertificateService.GetCertificateByID(certificate.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (cc *CertificateController) GetAllCertificates(c *gin.Context) {
	userID := c.GetUint("user_id")
	var query models.CertificateQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	certificates, err := cc.CertificateService.GetAllCertificates(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, certificates)
}
