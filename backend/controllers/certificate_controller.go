package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CertificateController interface {
	CreateCertificate(c *gin.Context)
	UpdateCertificate(c *gin.Context)
	DeleteCertificate(c *gin.Context)
	GetCertificateByID(c *gin.Context)
	GetAllCertificates(c *gin.Context)
}

type certificateController struct {
	certificateService services.CertificateService
}

func NewCertificateController(certificateService services.CertificateService) CertificateController {
	return &certificateController{certificateService: certificateService}
}

func (cc certificateController) CreateCertificate(c *gin.Context) {
	var req models.CertificateCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	certificate := models.Certificate{
		UserID:         req.UserID,
		Name:           req.Name,
		Organization:   req.Organization,
		CredentialID:   req.CredentialID,
		CredentialURL:  req.CredentialURL,
		IssueDate:      req.IssueDate,
		ExpirationDate: req.ExpirationDate,
		Image:          req.Image,
	}

	if err := cc.certificateService.CreateCertificate(&certificate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificate created successfully"})
}

func (cc certificateController) UpdateCertificate(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.CertificateUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	certificate, err := cc.certificateService.GetCertificateByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	certificate.Name = req.Name
	certificate.Organization = req.Organization
	certificate.CredentialID = req.CredentialID
	certificate.CredentialURL = req.CredentialURL
	certificate.IssueDate = req.IssueDate
	certificate.ExpirationDate = req.ExpirationDate
	certificate.Image = req.Image

	if err := cc.certificateService.UpdateCertificate(certificate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificate updated successfully"})
}

func (cc certificateController) DeleteCertificate(c *gin.Context) {
	var req models.CertificateDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	certificate, err := cc.certificateService.GetCertificateByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificate not found"})
		return
	}

	if err := cc.certificateService.DeleteCertificate(certificate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificate deleted successfully"})
}

func (cc certificateController) GetCertificateByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := cc.certificateService.GetCertificateByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (cc certificateController) GetAllCertificates(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var query models.CertificateQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Page == 0 {
		query.Page = 1
	}
	if query.Limit == 0 {
		query.Limit = 10
	}

	certificates, total, err := cc.certificateService.GetAllCertificates(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": certificates, "total": total})
}
