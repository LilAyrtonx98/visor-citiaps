<template>

  <div class="map-container">
      <div v-show="showFullLegend" style="position: absolute; top: 0; left: 0; z-index: -1;" ref="fullLegend" id="fullLegendID"> 
        <div
        v-for="(layer,i) in activeLayers" v-bind:key="i"
        >
          <strong>{{layer.name}}</strong>
          <img :src="getLegendSource(layer)" width="100%"> 
        </div>
      </div>

    <div ref="printMe" id="map" class="map">
      <div style="display: none;">
      </div>
    </div> 
    
  </div>

</template>

<style src="ol/ol.css"></style>

<style>

</style>

<script>
// OpenLayers
import Map from 'ol/Map';
import View from 'ol/View';
import {Tile as TileLayer, Vector as VectorLayer} from 'ol/layer.js';import XYZ from 'ol/source/XYZ';
import {OSM, Vector as VectorSource} from 'ol/source.js';
import Overlay from 'ol/Overlay';
import {fromLonLat, toLonLat} from 'ol/proj.js';
import {defaults as defaultControls, Control, FullScreen, ZoomSlider, ScaleLine, Zoom} from 'ol/control.js';
import {inherits} from 'ol/util.js';
import Feature from 'ol/Feature.js';
import Geolocation from 'ol/Geolocation.js';
import Point from 'ol/geom/Point.js';
import {Circle as CircleStyle, Fill, Stroke, Style, Icon} from 'ol/style.js';
import {defaults as defaultInteractions, DragRotateAndZoom} from 'ol/interaction.js';
import Geometry from 'ol/geom/Geometry';

// Mixin generales
import {formatter} from '@/components/mixin/formatter.js'
import {vuex} from '@/components/mixin/vuex.js'

// Navegación y controles del mapa
import {mapControlsCustom} from '@/components/mapControls/mapControlsCustom.js'
import {mapControlsOpenLayers} from '@/components/mapControls/mapControlsOpenLayers.js'

// Mixin herramientas del mapa
import {selectPoint} from '@/components/mapTools/select/selectPoint.js'
import {selectArea} from '@/components/mapTools/select/selectArea.js'
import {layers} from '@/components/mapTools/layers/layers.js'
import {charts} from '@/components/mapTools/charts/charts.js'
import {measureArea} from '@/components/mapTools/measures/measureArea.js'
import {measureDistance} from '@/components/mapTools/measures/measureDistance.js'
import {measureRadius} from '@/components/mapTools/measures/measureRadius.js'
import {geoprocessingsInteraction} from '@/components/mapTools/geoprocessings/geoprocessingsInteraction.js'
import {geoprocessingsRequests} from '@/components/mapTools/geoprocessings/geoprocessingsRequests.js'
import {search} from '@/components/mapTools/search/search.js'
import {legends} from '@/components/mapControls/legends/legends.js'
import {exportMap} from '@/components/mapTools/export/exportMap.js'
import {changeBasemap} from '@/components/mapTools/baseMap/changeBasemap.js'
import {annotations} from '@/components/mapTools/annotations/annotations.js'

export default{
  props:[],
  mixins: [
    formatter, 
    mapControlsCustom, 
    mapControlsOpenLayers,
    selectPoint,
    selectArea,
    layers,
    charts,
    measureArea,
    measureDistance,
    measureRadius,
    geoprocessingsInteraction,
    geoprocessingsRequests,
    search,
    legends,
    exportMap,
    changeBasemap,
    annotations,
    vuex
    ],
  data(){
    return {
        map: null,
        geolocation: null,
        selectedMeasureOption: 'distance',
        label: null,
        showFullLegend: false,
    }
  },
  computed: {
    
  },
  methods: {
    initMap: function() {
        var base = new TileLayer({
              source: new OSM()
            });
        var map = new Map({
              target: 'map',
              controls: new defaultControls({
                zoom: false,
                rotate: false,
                attribution: false,
              }),
              layers: [
                  this.getLocalBasemap(), 
                  this.getOpenStreetMapBasemap()
              ],
              view: new View({
                  center: fromLonLat([-70.4000, -33.2700]),
                  zoom: 3.5,
                  minZoom: 0,
              })
            });
            this.map = map;

            this.changeBasemap('openStreetMap');
    },
    shortkeysHandler: function(shortkey) {

      switch (shortkey) {
        case 'zoomIn':
          this.map.getView().animate({
                zoom: this.map.getView().getZoom() + 1,
                duration: 100
            });
          break;
        case 'zoomOut':
          this.map.getView().animate({
                zoom: this.map.getView().getZoom() - 1,
                duration: 100
            });
          break;
        default:
          return;
      }

    },
    initGeolocation: function () {
      var geolocation = new Geolocation({
        // enableHighAccuracy must be set to true to have the heading value.
        trackingOptions: {
          enableHighAccuracy: true
        },
        projection: this.map.getView().getProjection()
      });
      // update the HTML page when the position changes.
      geolocation.on('change', function() {
        console.log('Position changed')
      });

      // handle geolocation error.
      geolocation.on('error', function(error) {
        console.log('geolocation error')
        alert(error.message)
      });

      var accuracyFeature = new Feature();
      geolocation.on('change:accuracyGeometry', function() {
        accuracyFeature.setGeometry(geolocation.getAccuracyGeometry());
      });

      var positionFeature = new Feature();
      positionFeature.setStyle(new Style({
        image: new CircleStyle({
          radius: 12,
          fill: new Fill({
            color: '#3399CC'
          }),
          stroke: new Stroke({
            color: '#fff',
            width: 4
          })
        })
      }));

      geolocation.on('change:position', function() {
        var coordinates = geolocation.getPosition();
        positionFeature.setGeometry(coordinates ?
          new Point(coordinates) : null);
      });

      new VectorLayer({
        map: this.map,
        source: new VectorSource({
          features: [accuracyFeature, positionFeature]
        })
      });

      this.geolocation = geolocation;
    },
    getGeolocation: function() {
      this.geolocation.setTracking(true);
    }
  },
  created(){
  
  },
  watch: {
    // Funciones a ejecutar cuando se selecciona una herramienta específica
    'mapToolSelected': function (newValue, oldValue) {
        // Remover interacciones de herramienta anterior en el mapa
        switch (oldValue) {
          case 'select':
            this.setListenerOnSelectPoint(false)
            this.setListenerOnSelectArea(false)
            // Restaurar cursor
            document.body.style.cursor = 'default'
            break;
          case 'chart':
            this.addOrRemoveChartLayerListener('');
            break;
          case 'measures':
            this.removeOnlyDistanceInteraction()
            this.removeOnlyAreaInteraction()
            this.removeOnlyRadiusInteraction()
            break;
          case 'geoprocessing':
            this.removeGeoInteraction();
            this.removeGeoprocessingPolygons();
            break;
          case 'annotations':
            document.body.style.cursor = 'default'
            this.removeAnnotationInteraction();
            this.removeAnnotationSelectorListener();
            this.removeAnnotationsByUserLayer();
            this.removeAnnotationsByGroupLayer();
          default:
            break;
        }

        // Agregar interacciones de herramienta nueva en el mapa
        switch (newValue) {
          case 'geoprocessing':
            //this.addGeoInteraction()
            break;
          default:
            break;
        }
    },
    // Funciones a ejecutar cuando se realiza alguna acción en una herramienta
    'mapToolExecuted': function (newValue, oldValue) {
        switch (newValue) {
          case 'select':
            if(this.mapToolValues('select').action == 'selectMode') {
              if(this.mapToolValues('select').data.mode == 'point') {
                this.setListenerOnSelectArea(false)
                this.setListenerOnSelectPoint(true)
                
                // Cambiar cursor
                document.body.style.cursor = 'pointer'
              } else if(this.mapToolValues('select').data.mode == 'area') {
                this.setListenerOnSelectPoint(false)
                // Quitar interacción de área anterior (si existe)
                this.setListenerOnSelectArea(false)
                this.setListenerOnSelectArea(true)
                // Cambiar cursor
                document.body.style.cursor = 'crosshair'
              }
            }
            break;
          case 'chart':
           if(this.mapToolValues('chart').action == 'addLayer') {
             this.addOrRemoveChartLayerListener(this.mapToolValues('chart').data.layer);
           }
            break;
          case 'annotations':
            if(this.mapToolValues('annotations').action=='addAnnotationSelector'){
              this.addAnnotationSelectorListener()
            }
            if(this.mapToolValues('annotations').action=='addListener'){
              document.body.style.cursor = 'crosshair'
              this.addAnnotationInteraction();
            }
            if(this.mapToolValues('annotations').action=='removeListener'){
              document.body.style.cursor = 'default'
              this.removeAnnotationInteraction();
            }
            if(this.mapToolValues('annotations').action=='showAnnotationsByUser'){
              this.getFeaturesByUser(this.mapToolValues('annotations').data.annotationsByUser);
            }
            if(this.mapToolValues('annotations').action=='showAnnotationsByGroup'){
              this.getFeaturesByGroup(this.mapToolValues('annotations').data.annotationsByGroup,this.mapToolValues('annotations').data.userId);
            }
            if(this.mapToolValues('annotations').action=='removeAnnotationsByUser'){
              this.removeAnnotationsByUserLayer()
            }
            if(this.mapToolValues('annotations').action=='removeAnnotationsByGroup'){
              this.removeAnnotationsByGroupLayer()
            }
            break;
          case 'measures':
            if(this.mapToolValues('measures').action =='selectMeasure') {
              if(this.mapToolValues('measures').data.mode == 'distance') {
                this.addMeasureDistanceInteraction();
                this.removeMeasureAreaInteraction(false);
                this.removeMeasureRadiusInteraction(false);
              } else if(this.mapToolValues('measures').data.mode == 'area') {
                this.addMeasureAreaInteraction();
                this.removeMeasureDistanceInteraction(false);
                this.removeMeasureRadiusInteraction(false);
              } else if(this.mapToolValues('measures').data.mode == 'radius') {
                this.addMeasureRadiusInteraction();
                this.removeMeasureAreaInteraction(false);
                this.removeMeasureDistanceInteraction(false);
              } 
            } else if(this.mapToolValues('measures').action == 'clean') {
              if(this.mapToolValues('measures').data.mode == 'distance') {
                this.cleanMeasureDistance();
              } else if(this.mapToolValues('measures').data.mode == 'area') {
                this.cleanMeasureArea();
              } else if(this.mapToolValues('measures').data.mode == 'radius') {
                this.cleanMeasureRadius();
              }
            } else if(this.mapToolValues('measures').action == 'visibilityChange') {
              if (this.mapToolValues('measures').data.mode == 'distance') {
                this.setMeasureDistanceVisibility(this.mapToolValues('measures').data.isVisible)
              }
              if (this.mapToolValues('measures').data.mode == 'area') {
                this.setMeasureAreaVisibility(this.mapToolValues('measures').data.isVisible)
              }
            }
            break;
          case 'search':
            if(this.mapToolValues('search').action == 'setLocation') {
              this.addLocationMarker(this.mapToolValues('search').data.lat, this.mapToolValues('search').data.lon);
            } else if(this.mapToolValues('search').action == 'removeLocation') {
              this.removeLoationMarker()
            }
            break;
          case 'export':
            if(this.mapToolValues('export').action == 'exportMap') {
              if(this.mapToolValues('export').data.format == 'image') {
                this.exportToPNG(this.mapToolValues('export').data.filename)
              } else if(this.mapToolValues('export').data.format == 'document') {
                this.exportToPDF(this.mapToolValues('export').data.filename)
              }
            }
          case 'basemap':
            if(this.mapToolValues('basemap').action == 'change') {
              this.changeBasemap(this.mapToolValues('basemap').data.mode);
            }
            break;
          case 'layers':
            if(this.mapToolValues('layers').action == 'addLayer') {
              //this.addLayerWMS(process.env.VUE_APP_GEOSERVER_URL + 'wms', this.mapToolValues('layers').data);
              this.addLayerWMS(this.mapToolValues('layers').data);
            } else if(this.mapToolValues('layers').action == 'removeLayer') {
              this.removeLayerWMS(this.mapToolValues('layers').data);
            }
            break;
          case 'geoprocessing':
            if(this.mapToolValues('geoprocessing').action == 'execute') {
              this.executeGeoprocessing(this.mapToolValues('geoprocessing').data.geoprocessingBody, this.mapToolValues('geoprocessing').data.geoprocessingUrl)
            } else if(this.mapToolValues('geoprocessing').action == 'clean') {
              this.cleanGeoDraw()
            } else if(this.mapToolValues('geoprocessing').action == 'cleanResults') {
              this.removeGeoprocessingPolygons()
            } else if (this.mapToolValues('geoprocessing').action == 'changeDrawType') {
              // Quitar interacción anterior
              this.removeGeoInteraction();
              // Limpiar polígonos dibujados
              this.removeGeoprocessingPolygons();
              // Agregar nueva interacción
              this.addGeoInteraction(this.mapToolValues('geoprocessing').data.drawType.code)
            }
            break;
          default:
            break;
        }
    }
  },
  mounted () {
    this.initMap();
    this.initGeolocation();
    this.addControlsCustom();
    this.addControlsOpenLayers();
  },
  updated() {
       
  }
}
</script>

<style lang="scss" scoped>

.map-container{
  background: rgb(3,68,12);  
  height: 100vh;
}

@media only screen and (max-width: 640px) {
  
}

</style>
