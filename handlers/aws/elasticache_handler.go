package aws

import (
	"net/http"
)

func (handler *AWSHandler) ElasticacheClustersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("elasticache")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCacheClusters(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "elasticache:DescribeCacheClusters is missing")
		} else {
			handler.cache.Set("elasticache", response)
			respondWithJSON(w, 200, response)
		}
	}
}
