import {unByKey} from 'ol/Observable.js';
import Overlay from 'ol/Overlay.js';
import {getArea, getLength} from 'ol/sphere.js';
import {LineString, Polygon, Circle} from 'ol/geom.js';
import Draw from 'ol/interaction/Draw.js';
import {Vector as VectorLayer} from 'ol/layer.js';
import { Vector as VectorSource} from 'ol/source.js';
import {Circle as CircleStyle, Fill, Stroke, Style} from 'ol/style.js';
import { Feature } from 'ol';

var tooltipOverlaysCircle = [];

var sourceCircle = new VectorSource();

var mousePosition

var vectorCircle = new VectorLayer({
  source: sourceCircle,
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

var isVectorCircle = false;

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
var helpTooltipCircle;


/**
 * The measure tooltip element.
 * @type {Element}
 */
var measureTooltipCircleElementCircle;


/**
 * Overlay to show the measurement.
 * @type {module:ol/Overlay}
 */
var measureTooltipCircle;


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


/**
 * Message to show when the user is drawing a circle.
 * @type {string}
 */
var continueCircleMsg = 'Haga click para terminar de dibujar';

var drawCircle;

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

var formatRadius = function(circle) {
  var radius = circle.getRadius();
  var output;
  if (radius > 100) {
    output = (Math.round(radius / 1000 * 100) / 100) +
        ' ' + 'km';
  } else {
    output = (Math.round(radius * 100) / 100) +
        ' ' + 'm';
  }
  return output;
};

function pointerMoveRadius (evt) {
  if (evt.dragging) {
    return;
  }

  // Actualizar posición del mouse
  mousePosition = evt.coordinate

  /** @type {string} */
  var helpMsg = 'Haga click para comenzar a dibujar';

  if (sketch) {
    var geom = (sketch.getGeometry());
    if (geom instanceof Polygon) {
      helpMsg = continuePolygonMsg;
    } else if (geom instanceof LineString) {
      helpMsg = continueLineMsg;
    } else if (geom instanceof Circle) {
      helpMsg = continueCircleMsg;
    }
  }

  helpTooltipElement.innerHTML = helpMsg;
  helpTooltipCircle.setPosition(evt.coordinate);

  helpTooltipElement.classList.remove('hidden');
  helpTooltipElement.style.display = 'inherit'
}

function mouseOutRadius() {
  helpTooltipElement.classList.add('hidden');
  helpTooltipElement.style.display = 'none'
}

export const measureRadius = {
    data() {
        return {
            message: 'Hola'
        }
    },
    methods: {
      addInteraction: function() {
        var that = this;
        var type = 'Circle';
        drawCircle = new Draw({
          source: sourceCircle,
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
        this.map.addInteraction(drawCircle);

        this.createMeasureTooltipCircle();
        this.createHelpTooltipCircle();

        var listener;
        drawCircle.on('drawstart',
          function(evt) {
            // set sketch
            sketch = evt.feature;

            /** @type {module:ol/coordinate~Coordinate|undefined} */
            var tooltipCoord = evt.coordinate;

            // Radio de la circunferencia
            var radiusLine

            listener = sketch.getGeometry().on('change', function(evt) {
              var geom = evt.target;
              var output;
              if (geom instanceof Polygon) {
                output = formatArea(geom);
                tooltipCoord = geom.getInteriorPoint().getCoordinates();
              } else if (geom instanceof LineString) {
                output = formatLength(geom);
                tooltipCoord = geom.getLastCoordinate();
              } else if(geom instanceof Circle) {
                output = formatRadius(geom)
                tooltipCoord = geom.getCenter()
                // Dibujar línea que va desde el centro del círculo a un extremo
                var startPoint = geom.getCenter()
                var endPoint = geom.getClosestPoint(geom.getLastCoordinate())

                // Se quita radio anterior
                //that.map.removeLayer(radiusLine)
                
                // Se dibuja radio nuevo
                /*
                radiusLine = new VectorLayer( {
                  source: new VectorSource({
                    features: [new Feature( {
                      geometry: new LineString([startPoint, endPoint])
                    })]
                  })
                })
                */

                // Quitar radio anterior
                //sourceCircle.getFeatures().forEach(feature => {
                  //sourceCircle.removeFeature(feature)
                //});

                if(radiusLine != undefined) {
                  sourceCircle.removeFeature(radiusLine)
                }
                
                radiusLine = new Feature( {
                  geometry: new LineString([startPoint, mousePosition])
                })

                // Agregar nuevo radio
                sourceCircle.addFeature(radiusLine)

                // Se agrega radio nuevo al mapa
                //that.map.addLayer(radiusLine)

              }
              measureTooltipCircleElementCircle.innerHTML = output;
              measureTooltipCircle.setPosition(tooltipCoord);
            });
          });

        drawCircle.on('drawend',
          function() {
            measureTooltipCircleElementCircle.className = 'tooltip tooltip-static';
            measureTooltipCircle.setOffset([0, -7]);
            // unset sketch
            sketch = null;
            // unset tooltip so that a new one can be created
            measureTooltipCircleElementCircle = null;
            that.createMeasureTooltipCircle();
            unByKey(listener);
          });
      },
      createHelpTooltipCircle: function() {
        if (helpTooltipElement) {
          helpTooltipElement.parentNode.removeChild(helpTooltipElement);
        }
        helpTooltipElement = document.createElement('div');
        helpTooltipElement.className = 'tooltip hidden';
        helpTooltipCircle = new Overlay({
          element: helpTooltipElement,
          offset: [15, 0],
          positioning: 'center-left'
        });
        helpTooltipCircle.set('type', 'radius');
        this.map.addOverlay(helpTooltipCircle);
      },
      createMeasureTooltipCircle: function() {
        if (measureTooltipCircleElementCircle) {
          measureTooltipCircleElementCircle.parentNode.removeChild(measureTooltipCircleElementCircle);
        }
        measureTooltipCircleElementCircle = document.createElement('div');
        measureTooltipCircleElementCircle.className = 'tooltip tooltip-measure';
        measureTooltipCircle = new Overlay({
          element: measureTooltipCircleElementCircle,
          offset: [0, -15],
          positioning: 'bottom-center'
        });
        measureTooltipCircle.set('type', 'radius');
        this.map.addOverlay(measureTooltipCircle);
        tooltipOverlaysCircle.push(measureTooltipCircle);
      },
      addMeasureRadiusInteraction: function() {
        if(!isVectorCircle) {
          this.map.addLayer(vectorCircle);
          isVectorCircle = true;
        }
        this.map.addEventListener('pointermove', pointerMoveRadius);
        this.map.getViewport().addEventListener('mouseout', mouseOutRadius);
        this.addInteraction();
        //Mostrar overlays existentes
        this.map.getOverlays().forEach(overlay => {
          if(overlay.get('type') == 'radius') {
            overlay.getElement().style.display = 'inherit';
          }
        });
      },
      removeMeasureRadiusInteraction: function(deleteDrawing) {

        if(!drawCircle) {
          return;
        }
        
        this.map.removeInteraction(drawCircle);

        drawCircle = true;

        this.map.removeEventListener('pointermove', pointerMoveRadius, false);
        this.map.getViewport().removeEventListener('mouseout', mouseOutRadius, false);

        if(deleteDrawing) {
          isVectorCircle = false;
          this.map.removeLayer(vectorCircle);

          for (let i = 0; i < tooltipOverlaysCircle.length; i++) {
            this.map.removeOverlay(tooltipOverlaysCircle[i]);
          }
          
          //Vaciar tooltips
          tooltipOverlaysCircle = [];
          //Vector nuevo
          sourceCircle = new VectorSource();
          vectorCircle = new VectorLayer({
            source: sourceCircle,
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
      removeOnlyRadiusInteraction: function() {
        this.map.removeInteraction(drawCircle);
        drawCircle = null;
        this.map.removeEventListener('pointermove', pointerMoveRadius, false);
        this.map.getViewport().removeEventListener('mouseout', mouseOutRadius, false);
      },
      cleanMeasureRadius: function() {
        this.removeMeasureRadiusInteraction(true);
        this.addMeasureRadiusInteraction();
      },
      setMeasureRadiusVisibility: function(isVisible) {
        if(isVisible) {
          this.map.addLayer(vectorCircle);
          //Ocultar overlays existentes
          this.map.getOverlays().forEach(overlay => {
            if(overlay.get('type') == 'radius') {
              overlay.getElement().style.display = 'inherit';
            }
          });
        } else {
          this.map.removeLayer(vectorCircle);
          //Ocultar overlays existentes
          this.map.getOverlays().forEach(overlay => {
            if(overlay.get('type') == 'radius') {
              overlay.getElement().style.display = 'none';
            }
          });

        }
      }
    }
  }