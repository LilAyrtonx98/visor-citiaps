import jsPDF from 'jspdf';
import {vuex} from '@/components/mixin/vuex.js'

export const exportMap = {
    data() {
        return {
            message: 'Hola',
            output: null,
            canvas: null,
        }
    },
    mixins:[vuex],
    methods: {
      async print() {
        const el = this.$refs.printMe;
        // add option type to get the image version
        // if not provided the promise will return 
        // the canvas.
        const options = {
          //type: 'dataURL'
        }
        this.output = await this.$html2canvas(el, options);

        this.addControlsCustom();
        this.addControlsOpenLayers();
      },
      exportToPNG: function (filename) { 
        if(filename=='') {
          filename = 'mapa.png'
        }
        if(!filename.includes('.png')) {
          filename = filename + '.png'
        }

        var that = this;
        that.selectMapControl('loadScreen', true)
        
        this.removeControls();

        this.print()
        .then( function(){
          that.selectMapControl('loadScreen', false)
          var link = document.createElement('a');
          link.href = that.output.toDataURL("image/png").replace("image/png", "image/octet-stream");
          link.download = filename;
          //Firefox requires the link to be in the body
          document.body.appendChild(link);
          //simulate click
          link.click();
          //remove the link when done
          document.body.removeChild(link);
        })
        .catch(function(){
          that.selectMapControl('loadScreen', false)
        })

      },
      exportToPDF: function (filename) {
        var that = this;

        if(filename == '') {
          filename = 'map';
        }
        that.selectMapControl('loadScreen', true)

        var dims = {
          a0: [1189, 841],
          a1: [841, 594],
          a2: [594, 420],
          a3: [420, 297],
          a4: [297, 210],
          a5: [210, 148],
          letter: [216, 279]
        };

        var format = 'letter';
        var resolution = 150; //dpi

        var dim = dims[format];
        var width = Math.round(dim[0] * resolution / 25.4);
        var height = Math.round(dim[1] * resolution / 25.4);
        var size = /** @type {module:ol/size~Size} */ (this.map.getSize());
        var extent = this.map.getView().calculateExtent(size);
        
      this.removeControls();

       this.print()
       .then(function(){
          var thatToo = that;

          var imgData = that.output.toDataURL(
            'image/png');          

          //portrait
          let scale = dim[0]/size[0]
          let margin = 15;
          let separation = 8;
          //landscape
          let scaleLandscape = dim[1]/size[0]
          //Posición de Y (línea) donde se está escribiendo
          //Si es imagen se le agrega también el largo de la imagen además de la separación
          let currentY = 0;
          
          if(size[0] < size[1]) {
            var doc = new jsPDF('p', 'mm', 'letter');
            while(size[1]*scale - margin*2 > 216) {
              scale = scale/1.1
            }
          } else {
            scale = scaleLandscape
            var doc = new jsPDF('l', 'mm', 'letter');
            while(size[0]*scale - margin*2 > 279) {
              scale = scale/1.1
            }
          }
          
          doc.setFontSize(24)
          doc.text(margin, margin, 'Visor de mapas')
          currentY+=margin; 

          doc.addImage(imgData, 'PNG', margin, currentY + separation, size[0]*scale - margin*2, size[1]*scale - margin*2);
          currentY+=separation + size[1]*scale - margin*2;

          doc.setFontSize(12)
          //doc.text(margin, currentY + separation, 'Generado por Usuario 1')
          //currentY+=separation;

          doc.text(margin, currentY + separation, 'Fecha: ' + that.dateFormatter(new Date()))
          currentY+=separation;
          
          //doc.text(margin, currentY + separation, 'Total de capas: ' + that.activeLayers.length)
          //currentY+=separation;
          
          that.showFullLegend = true;
          that.printLegend()
          .then(function(){
              thatToo.selectMapControl('loadScreen', false)

              var imgDataLegend = thatToo.legendImage.toDataURL(
              'image/png');
              let element = document.getElementById("fullLegendID")
              let weight = element.style.width;
              let height = element.style.height;

              doc.addPage()
              
              doc.addImage(imgDataLegend, 'PNG', margin, margin, weight, height);
              
              doc.save(filename + '.pdf');
              thatToo.showFullLegend = false;
          })
          .catch(function(){
            thatToo.selectMapControl('loadScreen', false)
            thatToo.showFullLegend = false;
          })
       })
       .catch(function(){
          that.selectMapControl('loadScreen', false)
          that.showFullLegend = false;
       })

      },
      printSomething: function() {
        console.log('Something')
      },
      exportMapLegends: function() {
        var that = this;
  
        var filename = 'legends.pdf'
  
        if(filename == '') {
          filename = 'map';
        }
        that.selectMapControl('loadScreen', true)
        that.showFullLegend = true;
        this.printLegend()
        .then(function(){
         that.selectMapControl('loadScreen', false)
          var imgData = that.legendImage.toDataURL(
            'image/png');
  
          var doc = new jsPDF('p', 'mm', 'letter');
          doc.setFontSize(20)
          doc.text('Paranyan loves jsPDF', 25, 35, {
            align: 'justify',
          })
          
          let element = document.getElementById("fullLegendID")
          let weight = element.style.width;
          let height = element.style.height;
        
          doc.addImage(imgData, 'PNG', 0, 40, weight, height);
          doc.save(filename + '.pdf');
          that.showFullLegend = fale;
        })
        .catch(function(){
          that.selectMapControl('loadScreen', false)
          that.showFullLegend = false;
        })
      },
      async printLegend() {
        const element = this.$refs.fullLegend;
        // add option type to get the image version
        // if not provided the promise will return 
        // the canvas.
        const opt = {
          //type: 'dataURL'
          allowTaint : false,
          logging: true,
          useCORS: true,
        }
        this.legendImage = await this.$html2canvas(element, opt);
      },
    }
  }