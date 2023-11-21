import {Tile as TileLayer, Vector as VectorLayer, Image as ImageLayer}  from 'ol/layer.js';
import OSM from 'ol/source/OSM';
import TileWMS from 'ol/source/TileWMS.js';
import VectorSource from 'ol/source/Vector';
import {ImageArcGISRest} from 'ol/source.js';
import {vuex} from '@/components/mixin/vuex.js'
import {bbox as bboxStrategy, tile as tileStrategy} from 'ol/loadingstrategy';
import GeoJSON from 'ol/format/GeoJSON';
import {createXYZ} from 'ol/tilegrid';
import {Circle as CircleStyle, Fill, Stroke, Style, Icon, Circle} from 'ol/style.js';
import EsriJSON from 'ol/format/EsriJSON';

export const layers = {
    data() {
        return {
            
        }
    },
    mixins:[vuex],
    methods: {
      
      addLayerWMS: function (layer) {
        if(layer.provider.name=="arcgis"){
          /* Argis layer configuration */
          var arcgisUrl = layer.provider.url; 
          console.log("url arcgis",layer.provider.url)
          var arcgisLayer = 
            new ImageLayer({
              source: new ImageArcGISRest({
                ratio: 1,
                params: {},
                url: arcgisUrl
              }),
              opacity: 1
            });

            // This show a scpecific layer format
            // /* Comentar desde aqui para ignorar vector, se puede profudizar en detalles de la capa */
            // var vectorSource = new VectorSource({
            //   format: new EsriJSON(),
            //   url: function(extent) {
            //     return arcgisUrl + '/0/query/?f=json&' +
            //         'returnGeometry=true&spatialRel=esriSpatialRelIntersects&geometry=' +
            //         encodeURIComponent('{"xmin":' + extent[0] + ',"ymin":' +
            //             extent[1] + ',"xmax":' + extent[2] + ',"ymax":' + extent[3] +
            //             ',"spatialReference":{"wkid":102100}}') +
            //         '&geometryType=esriGeometryEnvelope&inSR=102100&outFields=*' +
            //         '&outSR=102100';
            //   }
            // });
            /* Comentar hasta aqui*/
            // var arcgisLayer = new VectorLayer({
            //   source: vectorSource
            // }); 
            this.map.addLayer(arcgisLayer)
            layer.arcgisLayer=arcgisLayer
        }
        else {
          /* Geoserver layer configuration */
          /*var vectorSource = new VectorSource({
            format: new GeoJSON({ featureProjection: "EPSG:3857" }),
            url: layer.provider.url+'&srsname=EPSG:3857'
          }); 

          var vector = new VectorLayer({
            source: vectorSource,
          }) 

          this.map.addLayer(vector) */
          var url=layer.provider.url
          console.log("url geoserver",url)
          console.log("layer geoserver",layer)
          // Origen de la capa
          var wmsSource = new TileWMS({
            url: url,
            params: {
              'LAYERS': layer.provider.geoserverdata.workspace + ':' + layer.provider.geoserverdata.datastore,
              'TILED': true,
              'VERSION':'1.1.1',
              'LAYER_PROJ': '3857',
              /*
              'LAYER_WORK': layer.provider.geoserverdata.workspace,
              'LAYER_NAME': layer.provider.geoserverdata.filename*/
            },
            serverType: 'geoserver',
            crossOrigin: 'anonymous',
            //transition: 0,
            // Atributos personalizados
            queryLayers: layer.provider.geoserverdata.workspace + ':' + layer.provider.geoserverdata.filename,
          })

          // Capa formada para OpenLayers
          var tileLayer = new TileLayer({
            source: wmsSource,
            opacity: 1
          })

          // Agregar capa al mapa
          this.map.addLayer(tileLayer)

          // Agregar tileLayer como atributo de la capa
          layer.tileLayer = tileLayer
        }
        this.$store.dispatch('addActiveLayer', layer)
      },
      removeLayerWMS: function(layer) {
       for (let i = 0; i < this.activeLayers.length; i++) {
          if(this.activeLayers[i].id == layer.id) {

            // Quitar capa del mapa
            if(layer.provider.name=="arcgis") this.map.removeLayer(this.activeLayers[i].arcgisLayer)
            else this.map.removeLayer(this.activeLayers[i].tileLayer)

            // Quitar capa de layersWMS
            this.$store.dispatch('removeActiveLayer', layer)
          }
        }
      }
    }
  }