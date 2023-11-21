export const restLayer = {
    data() {
        return {
            layer: {
              desc: '',
              geoserver_data: {
                filename: '',
                coordinatessystem: 'EPSG:3857',
                workspace: '',
                datastore: ''
              },
              id_category: '-1',
              name: '',
              provider:{
                name: '',
                url: '',
                parsedurl: [],
                geoserverdata: {
                    filename: '',
                    coordinatessystem: 'EPSG:3857',
                    workspace: '',
                    datastore: '',
                    maxfeatures: '',
                    outputformat: '',
                    request: '',
                    service: '',
                    version: ''
                }
              }
            },
            layers: [],
        }
    },
    methods: {
      getLayer: function(){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer)
        .then(function (response) {
            that.layer = response.data;
            console.log('get Layer',that.layer);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getLayersPages: function(page){
        var that = this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/layers?page=' + page)
        .then(function (response) {
            that.layers = response.data;
            response.headers.totalpages == undefined ? that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      getLayersByCategory: function(page, idCategory){
        var that=this;
        this.$http.get(this.$store.getters.backendurl + 'api/v1/categories/'+ idCategory + '/layers?page=' + page)
        .then(function (response){
          that.layers = response.data;
          response.headers.totalpages == undefined ?  that.totalPages = 0 : that.totalPages = parseInt(response.headers.totalpages, 10);
        })
      },
      postLayer: function(){
        var that = this;
        this.$http.post(this.$store.getters.backendurl + 'api/v1/layers', 
        this.layer
        )
        .then(function (response) {
            //Obtener ID de la tarea recién creada y pasar a la vista de asignar usuarios
            let idLayer = response.data.id
            that.$router.push('/capas/'+ idLayer + '/asignar-mapas')
        })
        .catch(function (error) {
            console.log(error);
        });
      },
      putLayer: function(){
        var that = this;
        this.$http.put(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer,
        this.layer
        )
        .then(function (response) {
            that.$router.push('/capas/'+ that.$route.params.idLayer)
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
      deleteLayer: function(){
        var that = this;
        this.$http.delete(this.$store.getters.backendurl + 'api/v1/layers/' + this.$route.params.idLayer)
        .then(function (response) {
            that.$router.push('/capas')
        })
        .catch(function (error) {
            console.log(error)
            that.message = error.response.data.message;
        });
      },
    }
  }