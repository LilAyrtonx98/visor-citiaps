import {unByKey} from 'ol/Observable.js';
import Overlay from 'ol/Overlay.js';
import {getArea, getLength} from 'ol/sphere.js';
import {LineString, Polygon} from 'ol/geom.js';
import Draw from 'ol/interaction/Draw.js';
import {Vector as VectorLayer} from 'ol/layer.js';
import {Vector as VectorSource} from 'ol/source.js';
import {Circle as CircleStyle, Fill, Stroke, Style} from 'ol/style.js';

var tooltipOverlaysDistance = [];

var sourceDistance = new VectorSource();

var vectorDistance = new VectorLayer({
  source: sourceDistance,
  style: new Style({
    fill: new Fill({
      color: 'rgba(255, 255, 255, 0.2)'
    }),
    stroke: new Stroke({
      color: '#ffcc33',
      width: 2
    }),
    image: new CircleStyle({
      radius: 7,
      fill: new Fill({
        color: '#ffcc33'
      })
    })
  })
});

var isVectorDistance = false;

/**
 * Currently drawn feature.
 * @type {module:ol/Feature~Feature}
 */
var sketch;


/**
 * The help tooltip element.
 * @type {Element}
 */
var helpTooltipElement;


/**
 * Overlay to show the help messages.
 * @type {module:ol/Overlay}
 */
var helpTooltipDistance;


/**
 * The measure tooltip element.
 * @type {Element}
 */
var measureTooltipElement;


/**
 * Overlay to show the measurement.
 * @type {module:ol/Overlay}
 */
var measureTooltip;


/**
 * Message to show when the user is drawing a polygon.
 * @type {string}
 */
var continuePolygonMsg = 'Haga click para continuar dibujando el polígono';


/**
 * Message to show when the user is drawing a line.
 * @type {string}
 */
var continueLineMsg = 'Haga click para continuar dibujando la línea';

var typeSelect = document.getElementById('type');

var drawDistance;

var formatLength = function(line) {
  var length = getLength(line);
  var output;
  if (length > 100) {
    output = (Math.round(length / 1000 * 100) / 100) +
        ' ' + 'km';
  } else {
    output = (Math.round(length * 100) / 100) +
        ' ' + 'm';
  }
  return output;
};

var formatArea = function(polygon) {
  var area = getArea(polygon);
  var output;
  if (area > 10000) {
    output = (Math.round(area / 1000000 * 100) / 100) +
        ' ' + 'km<sup>2</sup>';
  } else {
    output = (Math.round(area * 100) / 100) +
        ' ' + 'm<sup>2</sup>';
  }
  return output;
};

function pointerMoveDistance(evt) {
  if (evt.dragging) {
    return;
  }
  /** @type {string} */
  var helpMsg = 'Haga click para comenzar a dibujar';

  if (sketch) {
    var geom = (sketch.getGeometry());
    if (geom instanceof Polygon) {
      helpMsg = continuePolygonMsg;
    } else if (geom instanceof LineString) {
      helpMsg = continueLineMsg;
    }
  }

  helpTooltipElement.innerHTML = helpMsg;
  helpTooltipDistance.setPosition(evt.coordinate);

  helpTooltipElement.classList.remove('hidden');
  helpTooltipElement.style.display = 'inherit'
}

function mouseOutDistane() {
  helpTooltipElement.classList.add('hidden');
  helpTooltipElement.style.display = 'none'
}

export const measureDistance = {
    data() {
        return {
            message: 'Hola'
        }
    },
    methods: {
      addDistanceInteraction: function() {
        var that = this;
        //var type = (typeSelect.value == 'area' ? 'Polygon' : 'LineString');
        var type = 'LineString';
        drawDistance = new Draw({
          source: sourceDistance,
          type: type,
          style: new Style({
            fill: new Fill({
              color: 'rgba(255, 255, 255, 0.2)'
            }),
            stroke: new Stroke({
              color: 'rgba(0, 0, 0, 0.5)',
              lineDash: [10, 10],
              width: 2
            }),
            image: new CircleStyle({
              radius: 5,
              stroke: new Stroke({
                color: 'rgba(0, 0, 0, 0.7)'
              }),
              fill: new Fill({
                color: 'rgba(255, 255, 255, 0.2)'
              })
            })
          })
        });
        this.map.addInteraction(drawDistance);

        this.createDistanceMeasureTooltip();
        this.createDistanceHelpTooltip();

        var listener;
        drawDistance.on('drawstart',
          function(evt) {
            // set sketch
            sketch = evt.feature;

            /** @type {module:ol/coordinate~Coordinate|undefined} */
            var tooltipCoord = evt.coordinate;

            listener = sketch.getGeometry().on('change', function(evt) {
              var geom = evt.target;
              var output;
              if (geom instanceof Polygon) {
                output = formatArea(geom);
                tooltipCoord = geom.getInteriorPoint().getCoordinates();
              } else if (geom instanceof LineString) {
                output = formatLength(geom);
                tooltipCoord = geom.getLastCoordinate();
              }
              measureTooltipElement.innerHTML = output;
              measureTooltip.setPosition(tooltipCoord);
            });
          });

        drawDistance.on('drawend',
          function() {
            measureTooltipElement.className = 'tooltip tooltip-static';
            measureTooltip.setOffset([0, -7]);
            // unset sketch
            sketch = null;
            // unset tooltip so that a new one can be created
            measureTooltipElement = null;
            that.createDistanceMeasureTooltip();
            unByKey(listener);
          });
      },
      createDistanceHelpTooltip: function() {
        if (helpTooltipElement) {
          helpTooltipElement.parentNode.removeChild(helpTooltipElement);
        }
        helpTooltipElement = document.createElement('div');
        helpTooltipElement.className = 'tooltip hidden';
        helpTooltipDistance = new Overlay({
          element: helpTooltipElement,
          offset: [15, 0],
          positioning: 'center-left',
          type: 'distance'
        });
        helpTooltipDistance.set('type', 'distance');
        this.map.addOverlay(helpTooltipDistance);
      },
      createDistanceMeasureTooltip: function() {
        if (measureTooltipElement) {
          measureTooltipElement.parentNode.removeChild(measureTooltipElement);
        }
        measureTooltipElement = document.createElement('div');
        measureTooltipElement.className = 'tooltip tooltip-measure';
        measureTooltip = new Overlay({
          element: measureTooltipElement,
          offset: [0, -15],
          positioning: 'bottom-center'
        });
        measureTooltip.set('type', 'distance');
        this.map.addOverlay(measureTooltip);
        tooltipOverlaysDistance.push(measureTooltip);
      },
      addMeasureDistanceInteraction: function() {
        if(!isVectorDistance) {
          this.map.addLayer(vectorDistance);
          isVectorDistance = true;
        }
        this.map.addEventListener('pointermove', pointerMoveDistance);
        this.map.getViewport().addEventListener('mouseout', mouseOutDistane);
        this.addDistanceInteraction();
        //Mostrar overlays existentes
        this.map.getOverlays().forEach(overlay => {
          if(overlay.get('type') == 'distance') {
            overlay.getElement().style.display = 'inherit';
          }
        });
      },
      removeMeasureDistanceInteraction: function(deleteDrawing) {
        
        if(!drawDistance) {
          return;
        }

        this.map.removeInteraction(drawDistance);

        drawDistance = true;
      
        this.map.removeEventListener('pointermove', pointerMoveDistance, false);
        this.map.getViewport().removeEventListener('mouseout', mouseOutDistane, false);
      
        if(deleteDrawing) {

          this.map.removeLayer(vectorDistance);
          isVectorDistance = false;

          for (let i = 0; i < tooltipOverlaysDistance.length; i++) {
            this.map.removeOverlay(tooltipOverlaysDistance[i]);
          }
          
          //Vaciar tooltips
          tooltipOverlaysDistance = [];
          //Vector nuevo
          sourceDistance = new VectorSource();
          vectorDistance = new VectorLayer({
            source: sourceDistance,
            style: new Style({
              fill: new Fill({
                color: 'rgba(255, 255, 255, 0.2)'
              }),
              stroke: new Stroke({
                color: '#ffcc33',
                width: 2
              }),
              image: new CircleStyle({
                radius: 7,
                fill: new Fill({
                  color: '#ffcc33'
                })
              })
            })
          });
        }
        
      },
      removeOnlyDistanceInteraction: function() {
        this.map.removeInteraction(drawDistance);
        drawDistance = null;
        this.map.removeEventListener('pointermove', pointerMoveDistance, false);
        this.map.getViewport().removeEventListener('mouseout', mouseOutDistane, false);
      },
      cleanMeasureDistance: function() {

        this.removeMeasureDistanceInteraction(true);
        this.addMeasureDistanceInteraction();

      },
      setMeasureDistanceVisibility: function(isVisible) {
        if(isVisible) {
          this.map.addLayer(vectorDistance);
          //Ocultar overlays existentes
          this.map.getOverlays().forEach(overlay => {
            console.log(overlay.get('type'))
            if(overlay.get('type') == 'distance') {
              overlay.getElement().style.display = 'inherit';
            }
          });
        } else {
          this.map.removeLayer(vectorDistance);
          //Ocultar overlays existentes
          this.map.getOverlays().forEach(overlay => {
            console.log(overlay.get('type'))
            if(overlay.get('type') == 'distance') {
              overlay.getElement().style.display = 'none';
            }
          });

        }
      }
    }
  }