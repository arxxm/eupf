package rest

import (
	"github.com/edgecomllc/eupf/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListPfcpAssociations godoc
// @Summary List PFCP associations
// @Description List PFCP associations
// @Tags PFCP
// @Produce  json
// @Success 200 {object} NodeAssociationMapNoSession
// @Router /pfcp_associations [get]
func (h *Handler) listPfcpAssociations(c *gin.Context) {

	nodeAssociationsNoSession := make(domain.NodeAssociationMapNoSession)
	for k, v := range h.PfcpSrv.NodeAssociations {
		nodeAssociationsNoSession[k] = domain.NodeAssociationNoSession{
			ID:            v.ID,
			Addr:          v.Addr,
			NextSessionID: v.NextSessionID,
		}
	}
	c.IndentedJSON(http.StatusOK, nodeAssociationsNoSession)
}

// ListPfcpAssociationsFull godoc
// @Summary List PFCP associations
// @Description List PFCP associations
// @Tags PFCP
// @Produce  json
// @Success 200 {object} map[string]core.NodeAssociation
// @Router /pfcp_associations/full [get]
func (h *Handler) listPfcpAssociationsFull(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.PfcpSrv.NodeAssociations)
}