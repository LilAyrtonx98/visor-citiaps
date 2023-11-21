<template>
    <modal name="showInfo" style="z-index: 100000001" :height="'auto'" :width="'90%'">
        <div class="modal-content" v-if="!deleteWarning">
            <div class="modal-item container-1">
                <strong v-if="userName!=''">Usuario:</strong>
                <p v-if="userName!=''">{{userName}}</p>
                <strong>Texto:</strong>
                <p v-if="!editMode">{{info.text}}</p>
                <input v-else v-model="newInfo.text">
                <template v-if="!editMode">
                <p v-if="userName==''"><strong>Anotación compartida: </strong><span v-if="info.shared">Sí</span><span v-if="!info.shared">No</span></p>
                </template>
                <template v-else>
                    <strong>Anotación compartida: </strong>
                    <select v-model="newInfo.shared">
                        <option :value="true">Sí</option>
                        <option :value="false">No</option>
                    </select>
                    <p style="color:red; cursor:pointer" @click="deleteWarning=true"> 
                        <font-awesome-layers class="fa-lg">
                            <font-awesome-icon icon="trash-alt"/>
                        </font-awesome-layers> 
                        <strong style="text-decoration: underline red"> Eliminar</strong>
                    </p>
                </template>
            </div>
            <div v-if="!editMode" class="buttons-container modal-buttons" >
              <button v-on:click="$modal.hide('showInfo')" class="btn btn-secondary">Cerrar</button>
              <button v-if="userName==''" v-on:click="editAnnotation()" class="btn btn-primary">Editar</button>
            </div>
            <div v-else class="buttons-container modal-buttons" >
              <button v-on:click="editMode=false" class="btn btn-secondary">Cancelar</button>
              <button v-if="userName==''" v-on:click="putAnnotation()" class="btn btn-primary">Guardar</button>
            </div>
        </div>
        <div v-else class="modal-content">
            <div class="modal-item container-1">
                ¿Está seguro que desea eliminar la anotación?
            </div>
            <div class="buttons-container modal-buttons" >
              <button v-on:click="deleteWarning=false" class="btn btn-secondary">No</button>
              <button v-if="userName==''" v-on:click="deleteAnnotation()" class="btn btn-primary">Sí</button>
            </div>
        </div>
    </modal>
</template>

<script>
export default {
    name: 'AnnotationInfo',
    props:['info'],
    data(){
        return{
            userName:"",
            editMode: false,
            newInfo:{
                text:"",
                shared: false
            },
            deleteWarning: false
        }
    },
    methods: {
        putAnnotation: function(){
            console.log("EDITAR")
            this.info.text=this.newInfo.text
            this.info.shared=this.newInfo.shared

            var newData = {
                "id_user": this.info.user,
                "id_map": this.info.map,
                "id_group": this.info.group,
                "text": this.info.text,
                "location": {
                    "type": "Point",
                    "coordinates": this.info.coordinates
                },
                "is_shared": this.info.shared
            }

            var that=this
            this.$http.put(this.$store.getters.backendurl + 'api/v1/annotations/' + this.info.id, newData)
            .then(function(response){
              that.editMode=false
              that.$emit('saveEdit')
            })
            .catch(function (error) {
                console.log(error)
            });
        },
        deleteAnnotation: function(){
            console.log("eliminar anotacion")
            var that=this
            this.$http.delete(this.$store.getters.backendurl + 'api/v1/annotations/' + this.info.id)
            .then(function(response){
              that.editMode=that.$modal.hide("showInfo")
              that.deleteWarning=false
              that.$emit('saveEdit')
            })
            .catch(function (error) {
                console.log(error)
            });
        },
        editAnnotation: function(){
            this.newInfo.text=this.info.text
            this.newInfo.shared=this.info.shared
            this.editMode=true
        },
        getAnnotationUser: function(userId){
            var that=this
            this.$http.get(this.$store.getters.backendurl + 'api/v1/users/' + userId)
            .then(function(response){
                that.userName=response.data.firstname+" "+response.data.lastname
                console.log(response.data)
            })
            .catch(function (error) {
                that.userName="Sin nombre"
                console.log(error)
            });
        },
        getMyself: function(){
            var that = this;
            this.$http.get(this.$store.getters.backendurl + 'api/v1/accounts/me')
            .then(function (response) {
                that.info.user=response.data.id;
            })
            .catch(function (error) {
                console.log(error)
            });
        }
    },
    watch:{
        info: function(value){
            if(value){
                console.log("info", value)
                this.userName=""
                this.$modal.show("showInfo")
                if(value.user) this.getAnnotationUser(value.user)
                else this.getMyself()
            }
        }
    }
}
</script>