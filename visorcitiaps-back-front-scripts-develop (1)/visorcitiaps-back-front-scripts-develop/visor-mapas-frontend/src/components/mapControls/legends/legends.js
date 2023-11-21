export const legends = {
    data() {
        return {
          //activeLayers: []
        }
    },
    methods: {

      getLegendSource: function(layer){
        if(layer.provider.name == 'geoserver'){
          console.log('GEOSERVER getLegendSource');
          var url = layer.provider.parsed_url.host +':'+layer.provider.parsed_url.port + '/geoserver/wms?'
          var params =
          '&REQUEST=GetLegendGraphic' +
          '&VERSION=1.1.1' + 
          '&FORMAT=image/png' + 
          '&WIDTH=100' +
          '&HEIGHT=20' + 
          '&LAYER=' +  layer.provider.geoserverdata.workspace + ':' + layer.provider.geoserverdata.datastore +
          '&LAYER_WORK=' + layer.provider.geoserverdata.workspace
          return url+params
        }else{
          console.log('ARGIS getLegendSource');
        }
      }
    }
  }