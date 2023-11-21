import {Control} from 'ol/control.js';
import WMSCapabilities from 'ol/format/WMSCapabilities.js';
import {vuex} from '@/components/mixin/vuex.js'

//Barra de controles
var element = null
var element2 = null

export const mapControlsCustom = {
    data() {
        return {
          
        }
    },
    mixins:[vuex],
    methods: {
      addCustom: function () {
        //Custom control
        var button = document.createElement('button');
        button.className = 'icon custom';
        button.type = 'button';
        
        element.appendChild(button);

        var CustomControl = function() {
            console.log('Evento')
        };

        button.addEventListener('click', CustomControl, false);

        var CustomControl = new Control({
            element: element
        });

        this.map.addControl(CustomControl);
      },
      addZoomIn: function () {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon zoom-in';
        button.type = 'button';

        element.appendChild(button);

        var ZoomInControl = function() {
            var view = that.map.getView();
            view.animate({
                zoom: view.getZoom() + 1,
                duration: 100
            });
        };

        button.addEventListener('click', ZoomInControl, false);

        var ZoomInControl = new Control({
            element: element
        });

        this.map.addControl(ZoomInControl);
      },
      addZoomOut: function () {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon zoom-out';
        button.type = 'button';

        element.appendChild(button);

        var ZoomOutControl = function() {
            var view = that.map.getView();
            view.animate({
                zoom: view.getZoom() - 1,
                duration: 100
            });
        };

        button.addEventListener('click', ZoomOutControl, false);

        var ZoomOutControl = new Control({
            element: element
        });

        this.map.addControl(ZoomOutControl);
      },
      addGeolocation: function() {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon geolocation'
        button.type = 'button';

        element.appendChild(button);

        var GeolocationControl = function() {
          that.geolocation.setTracking(true);
        };

        button.addEventListener('click', GeolocationControl, false);

        var GeolocationControl = new Control({
            element: element
        });

        this.map.addControl(GeolocationControl);
      },
      addFullscreen: function() {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon full-screen'
        button.type = 'button';
        
        var FullScreenControl = function() {
          that.$emit('onClickFullScreen');
        };

        button.addEventListener('click', FullScreenControl, false);

        element.appendChild(button);

        var FullScreenControl = new Control({
            element: element
        });

        this.map.addControl(FullScreenControl);
      },
      addDrawer: function() {
        var that = this;
        
        var element = document.createElement('div');
        element.className = 'drawer-control ol-unselectable ol-control';

        var button = document.createElement('button');
        button.className = 'icon toolbox'
        button.type = 'button';

        element.appendChild(button);
        
        var DrawerControl = function() {
          that.selectMapControl('navDrawer', true)
        };

        button.addEventListener('click', DrawerControl, false);

        var DrawerControl = new Control({
            element: element
        });

        this.map.addControl(DrawerControl);
      },
      addHome: function () {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon home';
        button.type = 'button';
        
        element2.appendChild(button);

        var HomeControl = function() {
            that.$router.push('/inicio')
        };

        button.addEventListener('click', HomeControl, false);

        var HomeControl = new Control({
            element: element2
        });

        this.map.addControl(HomeControl);
      },
      addResetRotation: function () {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon reset-rotation';
        button.type = 'button';
        button.id = 'reset-rotation';
        button.style.display = 'none';
        
        element2.appendChild(button);

        this.map.getView().on('change:rotation', () => {
            let btn = document.getElementById("reset-rotation");
            if(that.map.getView().getRotation() != 0) {
              btn.style.display = 'inherit';
            } else {
              btn.style.display = 'none';
            }
        });

        var addResetRotation = function() {
            that.map.getView().animate({
              rotation : 0,
              duration: 200
            })
        };

        button.addEventListener('click', addResetRotation, false);

        var addResetRotation = new Control({
            element: element2
        });

        this.map.addControl(addResetRotation);
      },
      addLegend: function () {
        var that = this;

        var button = document.createElement('button');
        button.className = 'icon legend';
        button.type = 'button';
        
        element.appendChild(button);

        var LegendControl = function() {
            that.$emit('onClickMapLegend');
        };

        button.addEventListener('click', LegendControl, false);

        var LegendControl = new Control({
            element: element
        });

        this.map.addControl(LegendControl);
      },
      addShowCapabilities: function () {
        var parser = new WMSCapabilities();

        //Custom control
        var button = document.createElement('button');
        button.className = 'icon custom';
        button.type = 'button';
        
        element.appendChild(button);

        var CustomControl = function() {
            fetch('http://localhost:8600/geoserver/wms?service=WMS&version=1.1.0&request=GetCapabilities').then(function(response) {
              return response.text();
            }).then(function(text) {
              var result = parser.read(text);
            });
        };

        button.addEventListener('click', CustomControl, false);

        var CustomControl = new Control({
            element: element
        });

        this.map.addControl(CustomControl);
      },
      addControlsCustom: function() {
        element = document.createElement('div');
        element.className = 'custom-control ol-unselectable ol-control';

        element2 = document.createElement('div');
        element2.className = 'custom-control-2 ol-unselectable ol-control';

        this.addZoomIn();
        this.addZoomOut();
        this.addLegend();
        this.addGeolocation();
        this.addFullscreen();

        //this.addHome();
        this.addResetRotation();

        this.addDrawer();
      },
      removeControls: function() {
        let auxControls = this.map.getControls().array_.slice();
        auxControls.forEach(control => {
            this.map.removeControl(control);
        });
      }
    }
  }