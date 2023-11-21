import VectorLayer from "ol/layer/Vector"
import VectorSource from "ol/source/Vector"
import {Circle as CircleStyle,Fill, Stroke, Style, Text} from 'ol/style.js';
import Feature from 'ol/Feature';
import {Point} from 'ol/geom';
import 'ol/ol.css';
import * as proj from 'ol/proj';

var annotationsByUser;
var annotationsByGroup;
var vectorAnnotationsByUser;
var vectorAnnotationsByGroup;

var styleUserAnnotatios = new Style({
    image: new CircleStyle({
      radius: 5,
      fill: new Fill({color: '#3399ff'}),
      stroke: new Stroke({color: 'white', width: 2})
    })
});

var styleGroupAnnotatios = new Style({
    image: new CircleStyle({
      radius: 5,
      fill: new Fill({color: '#ffbb33'}),
      stroke: new Stroke({color: 'white', width: 2})
    })
});


export const annotations = {
    data(){
        return{
            message: "holi",
            featuresByUser: [],
            featuresByGroup: []
        }
    },
    methods: {
        addAnnotationInteraction: function(){
            this.map.addEventListener('singleclick', this.showAnnotationModal)
            this.map.removeEventListener('pointermove', this.changeCursor)
        },
        showAnnotationModal: function(evt){
            this.executeMapToolAction('annotations', 'showModal', {annotationCoor: proj.transform(evt.coordinate, 'EPSG:3857', 'EPSG:4326')})
        },
        removeAnnotationInteraction: function(){
            this.map.on('pointermove', this.changeCursor)
            this.map.removeEventListener('singleclick',this.showAnnotationModal)
        },
        removeAnnotationSelectorListener: function(){
            this.map.removeEventListener('pointermove', this.changeCursor)
            this.map.removeEventListener('click',this.sendAnnotationInfo)
        },
        addAnnotationSelectorListener: function(){
            this.map.on('click', this.sendAnnotationInfo)
            this.map.on('pointermove', this.changeCursor)
        },
        changeCursor: function(evt){
            var feature = this.map.forEachFeatureAtPixel(evt.pixel,
                function(feature) {
                return feature;
            });
            if(feature) document.body.style.cursor = 'pointer'
            else document.body.style.cursor = 'default'
        },
        sendAnnotationInfo: function(evt){
            var feature = this.map.forEachFeatureAtPixel(evt.pixel,
                function(feature) {
                return feature;
            });

            var annotationInfo = {
                id: "",
                text: "",
                user: "",
                shared: "",
                map: "",
                group: "",
                coordinates: [],
            }
            if(feature){
                document.body.style.cursor = 'default'
                annotationInfo.id=feature.get('id')
                annotationInfo.text=feature.get('text')
                annotationInfo.shared=feature.get('shared')
                annotationInfo.map=feature.get('map')
                annotationInfo.group=feature.get('group')
                annotationInfo.coordinates=feature.get('coordinates')
                if(feature.get('user')){
                    annotationInfo.user=feature.get('user')
                    this.executeMapToolAction('annotations','showAnnotationInfo',{info: annotationInfo})
                }
                else{
                    this.executeMapToolAction('annotations','showAnnotationInfo',{info: annotationInfo})
                }
            }
        },
        getFeaturesByUser: function(array){
            annotationsByUser=array
            this.featuresByUser = new Array(annotationsByUser.length);
            for (var i = 0; i < annotationsByUser.length; ++i) {
                this.featuresByUser[i] = new Feature({
                    'geometry': new Point(proj.fromLonLat(annotationsByUser[i].location.coordinates)),
                    'id': annotationsByUser[i].id,
                    'text': annotationsByUser[i].text,
                    'map': annotationsByUser[i].id_map,
                    'group': annotationsByUser[i].id_group,
                    'coordinates': annotationsByUser[i].location.coordinates,
                    'shared': annotationsByUser[i].is_shared
                });
            }
        },
        createAnnatotionsByUserLayer: function(features){
            vectorAnnotationsByUser = new VectorLayer({
                source: new VectorSource({
                    features: features,
                }),
                style: function() {
                    return styleUserAnnotatios;
                }
            });
            this.map.addLayer(vectorAnnotationsByUser); 

        },
        removeAnnotationsByUserLayer: function(){
            this.map.removeLayer(vectorAnnotationsByUser)
            vectorAnnotationsByUser=null
        },
        getFeaturesByGroup: function(array, user){
            annotationsByGroup=array
            this.featuresByGroup = new Array(annotationsByGroup.length);
            for(var i = 0; i < annotationsByGroup.length; ++i){
                if(annotationsByGroup[i].id_user!=user){
                    this.featuresByGroup[i] = new Feature({
                        'geometry': new Point(proj.fromLonLat(annotationsByGroup[i].location.coordinates)),
                        'id': annotationsByGroup[i].id,
                        'text': annotationsByGroup[i].text,
                        'map': annotationsByGroup[i].id_map,
                        'group': annotationsByGroup[i].id_group,
                        'coordinates': annotationsByGroup[i].location.coordinates,
                        'user': annotationsByGroup[i].id_user,
                        'shared': annotationsByGroup[i].is_shared
                    })
                }
            }
        },
        createAnnotationsByGroupLayer: function(features){
            vectorAnnotationsByGroup = new VectorLayer({
                source: new VectorSource({
                    features: features,
                }),
                style: function(){
                    return styleGroupAnnotatios;
                }
            });
            this.map.addLayer(vectorAnnotationsByGroup)
        },
        removeAnnotationsByGroupLayer: function(){
            this.map.removeLayer(vectorAnnotationsByGroup)
            vectorAnnotationsByGroup=null
        }
    },
    watch:{
        'featuresByUser': function(value){
            this.createAnnatotionsByUserLayer(value)
        },
        'featuresByGroup': function(value){
            var filtered = value.filter(function (el) {
                return el != null;
            });
            this.createAnnotationsByGroupLayer(filtered)
        }
    },
}