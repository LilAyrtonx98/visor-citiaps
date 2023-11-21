export const restVisor = {
    data() {
        return {
            mapsAssignedToUserID: []
        }
    },
    methods: {
        getMapsAssigned: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/visor/maps')
            .then(function (response) {
                that.mapsAssignedToUserID = response.data;
                console.log('maps', that.mapsAssignedToUserID)
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getLayersAssigned: function(idMap){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/visor/maps/' + idMap + '/layers' )
            .then(function (response) {
                console.log("CHAO", response)
                that.$store.dispatch('setLayers', response.data)
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        },
        getGeoprocessingsAssigned: function(idMap){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/visor/maps/' + idMap + '/geoprocessings' )
            .then(function (response) {
                that.$store.dispatch('setGeoprocessings', response.data)
            })
            .catch(function (error) {
                console.log(error)
                that.message = error.response.data.message;
            });
        }
    }
}