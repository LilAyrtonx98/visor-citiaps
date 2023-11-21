import Overlay from 'ol/Overlay';

var overlaysWMS = [];
var layersData= []

//Valor zIndex global
//Si hay dos popups, el que tenga mayor zIndex aparecerá delante del otro
var zIndex = 0;

function convertRemToPixels(rem) {    
  return rem * parseFloat(getComputedStyle(document.documentElement).fontSize);
}

export const selectPoint = {
    data() {
        return {
            message: 'Hola'
        }
    },
    methods: {

      // Crea el cuadro del popup vacío
      createEmptyPopup: function() {
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
      addContentPopup: function(containerPopup, headerText, count, content) {

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
      getLayerFeatureInfo: function(evt) {
            
        //Crear popup vacío
        var containerPopup = this.createEmptyPopup();
        var coordinate = evt.coordinate;

        //Agregar overlay (popup) al mapa
        containerPopup.overlay.setPosition(coordinate);
        this.map.addOverlay(containerPopup.overlay);
        
        var view = this.map.getView();
        var viewResolution = (view.getResolution());

        var that = this;

        // Reiniciar capas
        layersData = []

        var pixel = this.map.getEventPixel(evt.originalEvent);
        
        var baseSource = null
        this.map.forEachLayerAtPixel(pixel, function(layer) {
          var source = layer.getSource();
          // Si es capa WMS tiene la función getFeatureInfoUrl()
          if (typeof source.getFeatureInfoUrl === 'function') {
            // Guardar source de la primera capa WMS
            baseSource == null? baseSource = source : false
            // Agregar todas las capas WMS a la lista
            console.log("source", source)
            let data = {
              layerName: source.params_.LAYER_NAME,
              queryLayers: source.params_.LAYERS,
              properties: null,
              response: null,
              // Formar URL de consulta
              url: source.getFeatureInfoUrl(
                evt.coordinate, viewResolution, 'EPSG:3857',
                {
                  'INFO_FORMAT': 'application/json'
                })
            }
            console.log("data url", data.url)
            console.log("data real", data)
            layersData.push(data)
          }
          if(source.url_){
            console.log("test1 layer",source.getUrl()+ '?request=GetFeatureInfo')
            console.log("test12 layer",layer)

            var reqpoint = new XMLHttpRequest();
            reqpoint.open('GET', source.getUrl()+ '?request=GetFeatureInfo', false);
            reqpoint.withCredentials = false;
            reqpoint.send();

            if (reqpoint.status == 200){
              console.log("Success prueba", JSON.parse(reqpoint.response))
            }else{
              console.log("Fail prueba", reqpoint);
            }

            // that.$http.get(source.getUrl()+ '?request=GetFeatureInfo')
            // .then(function(response){
            //   console.log("prueba",response)
            // })
            // console.log("source", source.getUrl())
          }            
        },
        {
          hitTolerance: 10, // pixeles
          layerFilter: function(layer) {
            return true
          }
        })
        console.log("layersData", layersData)
        // Obtener la información de TODAS las capas antes de dibujar el popup
        let promiseArray = layersData.map(data => {
          var reqpoint = new XMLHttpRequest();
          reqpoint.open('GET', data.url, false);
          reqpoint.withCredentials = false;
          reqpoint.send();

          if (reqpoint.status == 200){
            console.log("Success prueba", JSON.parse(reqpoint.response))
            return JSON.parse(reqpoint.response)
          }else{
            console.log("Fail prueba", reqpoint);
            return {"Success": false}
          }
        });

        
        var results = []
        console.log("reqPozo empty ", results)
        console.log("Promise ", promiseArray)

        for (let i = 0; i < promiseArray.length; i++) {
          if (promiseArray[i]) {
            results.push(promiseArray[i])
          }
        }
        console.log("reqPozo full ", results)
        console.log("results", results)
        
        // let temp = results.map(r => r.data);
        // console.log("tmp", temp)

        for (let i = 0; i < results.length; i++) {
          if(results[i].features.length > 0){
            layersData[i].properties = results[i].features[0].properties
          }
        }

        // Quitar capas sin información en el punto
        layersData = layersData.filter(data => data.properties != null)

        // Dibujar el popup
        for (let i = 0; i < layersData.length; i++) {
          let aux = i + 1
          that.addContentPopup(
            containerPopup, 
            layersData[i].layerName,
              ' (' + aux + '/' + layersData.length + ')', 
              layersData[i].properties)
        }
        // Mensaje por defecto
        if(layersData.length < 1) {
          that.addContentPopup(containerPopup, "No hay datos", "", "")
        }
        
      },
      // Agrega o quita la interacción de seleccionar un punto en el mapa
      setListenerOnSelectPoint: function(isSet) {        
          switch (isSet) {
            case true:
              this.map.addEventListener('singleclick', this.getLayerFeatureInfo)
              this.map.getTargetElement().style.cursor = 'pointer'
              break;
            case false:
              this.map.removeEventListener('singleclick', this.getLayerFeatureInfo)
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

    }
  }