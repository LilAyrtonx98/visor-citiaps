import GeoJSON from 'ol/format/GeoJSON.js';
import {Vector as vectorLayer} from 'ol/layer.js';
import {Vector as vectorSource} from 'ol/source.js';
import {Circle as CircleStyle, Fill, Stroke, Style} from 'ol/style.js';

import {vuex} from '@/components/mixin/vuex.js'

var vectorSourceGeo = null
var vectorLayerGeo = null

var image = new CircleStyle({
  radius: 5,
  fill: null,
  stroke: new Stroke({color: 'red', width: 1})
});

var styles = {
  'Point': new Style({
    image: image
  }),
  'LineString': new Style({
    stroke: new Stroke({
      color: 'green',
      width: 1
    })
  }),
  'MultiLineString': new Style({
    stroke: new Stroke({
      color: 'green',
      width: 1
    })
  }),
  'MultiPoint': new Style({
    image: image
  }),
  'MultiPolygon': new Style({
    stroke: new Stroke({
      color: 'yellow',
      width: 1
    }),
    fill: new Fill({
      color: 'rgba(255, 255, 0, 0.9)'
    })
  }),
  'Polygon': new Style({
    stroke: new Stroke({
      color: 'blue',
      lineDash: [4],
      width: 3
    }),
    fill: new Fill({
      color: 'rgba(0, 0, 255, 0.1)'
    })
  }),
  'GeometryCollection': new Style({
    stroke: new Stroke({
      color: 'magenta',
      width: 2
    }),
    fill: new Fill({
      color: 'magenta'
    }),
    image: new CircleStyle({
      radius: 10,
      fill: null,
      stroke: new Stroke({
        color: 'magenta'
      })
    })
  }),
  'Circle': new Style({
    stroke: new Stroke({
      color: 'red',
      width: 2
    }),
    fill: new Fill({
      color: 'rgba(255,0,0,0.2)'
    })
  })
};

var geojsonObject = {
  type: 'FeatureCollection',
  crs: {
    type: 'name',
    properties: {
      name: 'EPSG:3857'
    }
  },
  features: []
}

var styleFunction = function(feature) {
  return styles[feature.getGeometry().getType()];
};

export const geoprocessingsRequests = {
    data() {
        return {
            message: 'Hola',
            
        }
    },
    mixins:[vuex],
    methods: {
      sendGeoprocessRequest: function(geoprocessingBody, url){
        var that = this;
        that.selectMapControl('loadScreen', true)
        this.removeGeoprocessingPolygons()
        this.$http.post(url, 
        geoprocessingBody
        )
        .then(function (response) {

          // Enviar atributos
          that.executeMapToolAction('geoprocessing', 'sendAttributes', {attributes: response.data.attributes})

          that.executeMapToolAction('geoprocessing', 'sendError', {errorMessage: 'ok'})

          // Limpiar features
          geojsonObject.features = []

          response.data.polygons.forEach(polygon => {

            let feature = {
              type: 'Feature',
              geometry: {}
            }

            polygon.requiresTransform = true

            feature.geometry = polygon

            geojsonObject.features.push(feature)
          });
          that.drawGeoprocessingPolygons()
        })
        .catch(function (error) {
            that.selectMapControl('loadScreen', false)
        });
      },
      executeGeoprocessing: function(geoprocessingBody, url) {
          geoprocessingBody.geometry.coordinates = this.getPolygonCoords()
          if(geoprocessingBody.geometry.type == 'Circle') {
            geoprocessingBody.radius = this.getCircleRadius()
          }
          console.log(geoprocessingBody)
          if(geoprocessingBody.layers.length == 0) {
            this.executeMapToolAction('geoprocessing', 'sendError', {errorMessage: 'No hay capas seleccionadas'})
          } else if(!geoprocessingBody.geometry) {
            this.executeMapToolAction('geoprocessing', 'sendError', {errorMessage: 'No se dibujó el polígono'})
          } else if(geoprocessingBody.geometry.coordinates == null){
              this.executeMapToolAction('geoprocessing', 'sendError', {errorMessage: 'No se dibujó el polígono'})
          } else {
            this.executeMapToolAction('geoprocessing', 'sendError', {errorMessage: ''})
            this.sendGeoprocessRequest(geoprocessingBody, url)
          }
      },
      drawGeoprocessingPolygons: function() {        
        vectorSourceGeo = new vectorSource({
          features: (new GeoJSON()).readFeatures(geojsonObject)
        });
        vectorLayerGeo = new vectorLayer({
          renderMode: 'image',
          source: vectorSourceGeo,
          style: styleFunction,
        });
        
        this.map.addLayer(vectorLayerGeo)
        this.selectMapControl('loadScreen', false)
      },
      removeGeoprocessingPolygons: function() {
        if(vectorLayerGeo!=null) {
          this.map.removeLayer(vectorLayerGeo)
          vectorLayerGeo = null
          vectorSourceGeo = null
        }
      }
    }
  }