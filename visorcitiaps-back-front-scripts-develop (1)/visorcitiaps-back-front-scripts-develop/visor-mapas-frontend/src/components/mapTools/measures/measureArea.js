import {unByKey} from 'ol/Observable.js';
import Overlay from 'ol/Overlay.js';
import {getArea, getLength} from 'ol/sphere.js';
import {LineString, Polygon} from 'ol/geom.js';
import Draw from 'ol/interaction/Draw.js';
import {Vector as VectorLayer} from 'ol/layer.js';
import {Vector as VectorSource} from 'ol/source.js';
import {Circle as CircleStyle, Fill, Stroke, Style} from 'ol/style.js';

var tooltipOverlaysArea = [];

var sourceArea = new VectorSource();

var vectorArea = new VectorLayer({
  source: sourceArea,
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

var isVectorArea = false;

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
var helpTooltipArea;


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

var drawArea;

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

function pointerMoveArea (evt) {
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
  helpTooltipArea.setPosition(evt.coordinate);

  helpTooltipElement.classList.remove('hidden');
  helpTooltipElement.style.display = 'inherit'
}

function mouseOutArea() {
  helpTooltipElement.classList.add('hidden');
  helpTooltipElement.style.display = 'none'
}

export const measureArea = {
    data() {
        return {
            message: 'Hola'
        }
    },
    methods: {
      addAreaInteraction: function() {
        console.log('add interaction 2')
        var that = this;
        //var type = (typeSelect.value == 'area' ? 'Polygon' : 'LineString');
        var type = 'Polygon';
        drawArea = new Draw({
          source: sourceArea,
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
        this.map.addInteraction(drawArea);

        this.createMeasureTooltip();
        this.createHelpTooltip();

        var listener;
        drawArea.on('drawstart',
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

        drawArea.on('drawend',
          function() {
            measureTooltipElement.className = 'tooltip tooltip-static';
            measureTooltip.setOffset([0, -7]);
            // unset sketch
            sketch = null;
            // unset tooltip so that a new one can be created
            measureTooltipElement = null;
            that.createMeasureTooltip();
            unByKey(listener);
          });
      },
      createHelpTooltip: function() {
        if (helpTooltipElement) {
          helpTooltipElement.parentNode.removeChild(helpTooltipElement);
        }
        helpTooltipElement = document.createElement('div');
        helpTooltipElement.className = 'tooltip hidden';
        helpTooltipArea = new Overlay({
          element: helpTooltipElement,
          offset: [15, 0],
          positioning: 'center-left'
        });
        helpTooltipArea.set('type', 'area');
        this.map.addOverlay(helpTooltipArea);
      },
      createMeasureTooltip: function() {
        console.log('create createMeasureTooltip')
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
        measureTooltip.set('type', 'area');
        this.map.addOverlay(measureTooltip);
        tooltipOverlaysArea.push(measureTooltip);
      },
      addMeasureAreaInteraction: function() {
        if(!isVectorArea) {
          this.map.addLayer(vectorArea);
          isVectorArea = true;
        }
        this.map.addEventListener('pointermove', pointerMoveArea);
        this.map.getViewport().addEventListener('mouseout', mouseOutArea);
        this.addAreaInteraction();
        //Mostrar overlays existentes
        this.map.getOverlays().forEach(overlay => {
          if(overlay.get('type') == 'area') {
            overlay.getElement().style.display = 'inherit';
          }
        });
      },
      removeMeasureAreaInteraction: function(deleteDrawing) {

        if(!drawArea) {
          return;
        }
        
        this.map.removeInteraction(drawArea);

        drawArea = true;

        this.map.removeEventListener('pointermove', pointerMoveArea, false);
        this.map.getViewport().removeEventListener('mouseout', mouseOutArea, false);

        if(deleteDrawing) {
          isVectorArea = false;
          this.map.removeLayer(vectorArea);

          for (let i = 0; i < tooltipOverlaysArea.length; i++) {
            this.map.removeOverlay(tooltipOverlaysArea[i]);
          }
          
          //Vaciar tooltips
          tooltipOverlaysArea = [];
          //Vector nuevo
          sourceArea = new VectorSource();
          vectorArea = new VectorLayer({
            source: sourceArea,
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
      removeOnlyAreaInteraction: function() {
        this.map.removeInteraction(drawArea);
        drawArea = null;
        this.map.removeEventListener('pointermove', pointerMoveArea, false);
        this.map.getViewport().removeEventListener('mouseout', mouseOutArea, false);
      },
      cleanMeasureArea: function() {
        this.removeMeasureAreaInteraction(true);
        this.addMeasureAreaInteraction();
      },
      setMeasureAreaVisibility: function(isVisible) {
        if(isVisible) {
          this.map.addLayer(vectorArea);
          //Ocultar overlays existentes
          this.map.getOverlays().forEach(overlay => {
            console.log(overlay.get('type'))
            if(overlay.get('type') == 'area') {
              overlay.getElement().style.display = 'inherit';
            }
          });
        } else {
          this.map.removeLayer(vectorArea);
          //Ocultar overlays existentes
          this.map.getOverlays().forEach(overlay => {
            console.log(overlay.get('type'))
            if(overlay.get('type') == 'area') {
              overlay.getElement().style.display = 'none';
            }
          });

        }
      }
    }
  }