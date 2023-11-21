<template>

    <div v-if="type=='success'" class="badge badge-success">{{message}}</div>
    <div v-else-if="type=='warning'" class="badge badge-warning">{{message}}</div>
    <div v-else-if="type=='danger'" class="badge badge-danger">{{message}}</div>
    <div v-else-if="type=='info'" class="badge badge-info">{{message}}</div>
    <div v-else-if="type=='secondary'" class="badge badge-secondary">{{message}}</div>
    <div v-else class="badge badge-primary">{{message}}</div>

</template>

<style lang="scss" scoped>
.badge{
  border: 1px solid rgba(0,0,0,0.5);
  border-radius: 6px;
  font-size: 0.75em;
  text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;
  font-weight: bold;
  min-width: 5em;
  text-align: center;
  align-self: center;
}
.badge-success{
  border: solid 2px rgb(10, 100, 30);
  background-color: rgb(15, 165, 45);
  color: white;
}
.badge-warning{
  border: solid 2px rgb(240, 160, 15);
  background-color: rgb(245, 225, 30);
  color: white;
}
.badge-danger{
  border: solid 2px rgb(115, 15, 5);
  background-color: rgb(185, 20, 20);
  color: white;
}
.badge-info{
  border: solid 2px rgb(25, 80, 105);
  background-color: rgb(15, 120, 170);
  color: white;
}
.badge-primary{
  border: solid 2px rgb(40, 30, 100);
  background-color: rgb(70, 57, 182);
  color: white;
}
.badge-secondary{
  border: solid 2px rgb(70, 57, 182);
  background-color: white;
  color: rgb(70, 57, 182);
}
</style>


<script>
export default{
  props:['type', 'message'],
  data(){
    return {
        
    }
  },
  computed: {

  },
  methods: {

  },
  created(){
  
  },
  watch: {

  }
}
</script>