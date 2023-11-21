export const restGeoprocessing = {
    data() {
        return {
            geo: {
                name: '',
                desc: '',
                geo_url: ''
            },
            geoprocessings: []
        }
    },
    methods: {
      getGeoprocessing: function(){
          var that = this;
          this.$http.get(this.$store.getters.backendurl + 'api/v1/geoprocessings/' + this.$route.params.idGeo)
          .then(function (response) {
              that.geo = response.data;
          })
          .catch(function (error) {
              console.log(error)
              that.message = error.response.data.message;
          });
      },
      getGeoprocessingsPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/geoprocessings?page=' + page)
        .then(function (response) {
            that.geoprocessings = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      postGeoprocessing: function(){
        var that = this;
        this.$http.post(this.$store.getters.backendurl + 'api/v1/geoprocessings', 
        this.geo
        )
        .then(function (response) {
            //Obtener ID de la tarea recién creada y pasar a la vista de asignar usuarios
            let idGeoprocessing = response.data.id
            that.$router.push('/geoprocesos/'+ idGeoprocessing + '/asignar-mapas')
        })
        .catch(function (error) {
            console.log(error);
        });
        },
      putGeoprocessing: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/geoprocessings/' + this.$route.params.idGeo,
        this.geo
        )
        .then(function (response) {
            that.$router.push('/geoprocesos/'+ that.$route.params.idGeo)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
        },
      deleteGeoprocessing: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/geoprocessings/' + this.$route.params.idGeo)
        .then(function (response) {
            that.$router.push('/geoprocesos')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }