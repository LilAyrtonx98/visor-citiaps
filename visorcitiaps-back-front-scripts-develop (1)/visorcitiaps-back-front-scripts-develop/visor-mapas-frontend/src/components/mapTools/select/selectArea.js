import Overlay from 'ol/Overlay';
import Draw from 'ol/interaction/Draw.js';
import {Vector as VectorLayer} from 'ol/layer.js';
import {Vector as VectorSource} from 'ol/source.js';
import GeoJSON from 'ol/format/GeoJSON';
import {bbox as bboxStrategy} from 'ol/loadingstrategy';
import {Stroke, Style, Fill} from 'ol/style';
import * as Extent from 'ol/extent';

import {vuex} from '@/components/mixin/vuex.js'

var overlaysWMS = [];
var layersData= []

//Valor zIndex global
//Si hay dos popups, el que tenga mayor zIndex aparecerá delante del otro
var zIndex = 0;

var draw; // global so we can remove it later

var source = new VectorSource({wrapX: false});

var vector;

var center;

var geometry;

function convertRemToPixels(rem) {    
  return rem * parseFloat(getComputedStyle(document.documentElement).fontSize);
}

export const selectArea = {
    data() {
        return {
            areaCoords: [0,0],
            layersDataShared: layersData
        }
    },
    mixins:[vuex],
    methods: {

      // Crea el cuadro del popup vacío
      createEmptyPopupArea: function() {
          //Aumentar zIndex
          zIndex = zIndex + 1 
          //Crear popup info
          var container = document.createElement('div');
          container.id = 'popup'
          container.className = 'ol-popup'
          container.style.zIndex = zIndex.toString();
          
          var closer = document.createElement('a');
          closer.id = 'popup-closer'
          closer.className = 'ol-popup-closer'
          //closer.href = '#'
          closer.style = 'font-size: 1.5em;'

          //Crear popup page
          var page = document.createElement('div');
          page.id = 'popup-page' + zIndex.toString();
          page.style.height = '20rem'
          page.style.overflowY = 'hidden'

          var row = document.createElement('div');
          row.id = 'popup-row'
          row.className = 'ol-popup-row'

          var next = document.createElement('a');
          next.id = 'popup-next'
          next.className = 'popup-button'
          //next.href = '#'
          next.style = 'font-size: 1em; margin-right: 1em;'
          next.className = 'next popup-button'
          //next.innerHTML = 'NEXT' 

          var prev = document.createElement('a');
          prev.id = 'popup-prev'
          prev.className = 'popup-button'
          //prev.href = '#'
          prev.style = 'font-size: 1em; margin-right: 1em;'
          prev.className = 'prev popup-button'
          //prev.innerHTML = 'PREV' 

          row.appendChild(prev)
          row.appendChild(next)
          
          container.appendChild(closer)
          container.appendChild(page)
          container.appendChild(row)

          //Crear overlay
          var overlay = new Overlay({
            element: container,
            autoPan: true,
            autoPanAnimation: {
              duration: 250
            }
          });

          overlay.set('type', 'wms');
          overlay.set('zIndex', zIndex)
          
          overlaysWMS.push(overlay);

          //Función del botón closer
          var closerAction = function() {
              overlay.setPosition(undefined);
              closer.blur();
              return false;
          };

          //Función del botón closer
          var nextAction = function() {
            var elem = document.getElementById('popup-page' + overlay.get('zIndex'))
            var scroll = convertRemToPixels(parseInt(elem.style.height, 10));
            elem.scrollTop += scroll;
          };

          //Función del botón closer
          var prevAction = function() {
            var elem = document.getElementById('popup-page' + overlay.get('zIndex'));
            var scroll = convertRemToPixels(parseInt(elem.style.height, 10));
            elem.scrollTop -= scroll;
          };

          //Listeners
          closer.addEventListener('click', closerAction, false);
          prev.addEventListener('click', prevAction, false);
          next.addEventListener('click', nextAction, false);

          return {'overlay': overlay, 'container': container}
      },
      // Agregar páginas al popup
      addContentPopupArea: function(containerPopup, headerText, count, content) {

          var contentText = ''

          Object.keys(content).forEach(key => {
            contentText = contentText + key + ': ' + content[key] + '<br>'
          })

          var info = document.createElement('div');
          info.style.height = "20rem";

          var header = document.createElement('div');
          header.id = 'popup-header'
          header.className = 'ol-popup-header'
          header.innerHTML = '<h4 style="margin: 0.25em 0;">Información de capa'+ count +'</h4><p style="margin-bottom: 0.25em;">' + 
          headerText + 
          '</p>';

          var content = document.createElement('div');
          content.id = 'popup-content'
          if(contentText != "") {
            content.style = 'font-size: 0.8em; border: solid 2px rgba(3, 96, 22, 0.5);'
            content.innerHTML = '<div class="popup-scrollable"><span> <meta charset="UTF-8">' + contentText + '</span></div>'
          }

          info.appendChild(header)
          info.appendChild(content)

          containerPopup.container.getElementsByTagName("div")[0].appendChild(info)
      },
      // Obtiene la información de las capas en el punto seleccionado
      getLayerFeatureInfoArea: function(polyCenter, polyGeom) {
        //Crear popup vacío
        var containerPopup = this.createEmptyPopupArea();

        // Asignar coordenada donde se va a ubicar el popup en el mapa
        var coordinate = polyCenter

        //Agregar overlay (popup) al mapa
        containerPopup.overlay.setPosition(coordinate);
        this.map.addOverlay(containerPopup.overlay);
        
        var that = this;

        // Reiniciar capas
        layersData = []

        // Reiniciar contador de capas listas
        var layerReadyCounter = 0
      
        this.activeLayers.forEach(function(layer) {

          var vectorSource = new VectorSource({
            format: new GeoJSON(),
            url: function(extent) {
              console.log('area url source',layer.provider.url+'&srsname=EPSG:3857&' +
              'bbox=' + extent.join(',') + ',EPSG:3857')
              return layer.provider.url+'&srsname=EPSG:3857&' +
                  'bbox=' + extent.join(',') + ',EPSG:3857';
            },
            strategy: bboxStrategy
          });
          console.log('area vector source',vectorSource)
          var vector = new VectorLayer({
            source: vectorSource,
            style: new Style({
              stroke: new Stroke({
                color: 'rgba(0, 0, 255, 1.0)',
                width: 2
              }),
              fill: new Fill({
                color: '#000',
                width:10
              }),
            })
          });
          
          that.map.addLayer(vector)

          vectorSource.on('change', function(event) {
            if (vectorSource.getState() == 'ready') {
              // Cantidad de capas listas
              layerReadyCounter++
              var extent = polyGeom.getExtent()
              vectorSource.forEachFeatureIntersectingExtent(extent, function(feature) {
                let data = {
                  layerName: layer.provider.geoserverdata.filename,
                  properties: feature.getProperties(),
                }

                layersData.push(data)

               })

               // Si el numero de capas listas es igual a la cantidad de capas activas, todas las capas están listas
               // Cuando todas las capas estén listas, se dibuja el pop-up en el mapa
               if(layerReadyCounter == that.activeLayers.length) {
                 // Dibujar el popup
                for (let i = 0; i < layersData.length; i++) {
                  let aux = i + 1
                  that.addContentPopupArea(
                    containerPopup, 
                    layersData[i].layerName,
                    ' (' + aux + '/' + layersData.length + ')', 
                    layersData[i].properties)
                }

                // Mensaje por defecto
                if(layersData.length < 1) {
                  that.addContentPopupArea(containerPopup, "No hay datos", "", "")
                }
               }
              that.map.removeLayer(vector)
            }
          });

        })

      },
      addFeatureCallback: function(evt) {
        var feature = evt.feature
        center = Extent.getCenter(feature.getGeometry().getExtent())
        geometry = evt.feature.getGeometry()
        this.getLayerFeatureInfoArea(center, geometry)
      },
      addDrawAreaInteraction: function() {
        var that = this
       source.on('addfeature', this.addFeatureCallback)
        vector = new VectorLayer({
          source: source
        });
        this.map.addLayer(vector);
        draw = new Draw({
          source: source,
          type: this.mapToolValues('select').data.drawType
          //type: 'Polygon'
        });
        draw.on('drawstart', function(evt){
          source.clear() // Elimina las capas anteriores
        })
        draw.on('drawend', function(evt){
        })
        this.map.addInteraction(draw);
      },
      removeDrawAreaInteraction: function() {
        this.map.removeLayer(vector);
        this.map.removeInteraction(draw);
        source.clear();
        draw = null;
        vector = null;
        center = null;
        source.un('addfeature', this.addFeatureCallback)
      },
      cleanDrawArea: function() {
        source.clear();
        center = null;
      },
      getPolygonCoords: function() {
        return this.areaCoords;
      },
      getPolygonGeometry: function() {
        return geometry
      },
      // Agrega o quita la interacción de seleccionar un punto en el mapa
      setListenerOnSelectArea: function(isSet) {        
          switch (isSet) {
            case true:
              this.addDrawAreaInteraction()
              this.map.getTargetElement().style.cursor = 'pointer'
              break;
            case false:
              this.removeDrawAreaInteraction()
              this.map.getTargetElement().style.cursor = ''
              for (let i = 0; i < overlaysWMS.length; i++) {
                this.map.removeOverlay(overlaysWMS[i]);
              }
              overlaysWMS = [];
              zIndex = 0;
              break;
            default:
              break;
          }
          
        },

    },
    watch: {
      'layersDataShared': function (newValue, oldValue) {
        console.log('largo: ' + newValue.length)
        if(newValue.length == this.activeLayers.length) {
          //console.log('iguales')
        }
      }
    }
  }