<template>
    <div class="autocomplete" style="height: max-content; position: absolute; display: contents;">

      <input v-model="query" type="text" placeholder='Buscar'/>

      <div class="dropdown-list">
        
        <template v-if="query.length==0">
          <p>{{beforeSearchMessage}}</p>
        </template>

        <template v-else-if="error">
          <p>{{emptySearchResultsMessage}}</p>
        </template>

        <template v-else-if="query.length!=0 && options.length==0">
          <p>{{searchingMessage}}</p>
        </template>

        <template v-else>
            <div class="dropdown-option" v-for="o in reducedOptions" v-bind:key="o.index" @click="selectItem(o)">
                {{o[attribute]}}
            </div>
        </template>

      </div>

  </div>
</template>

<style lang="scss" scoped>

$main-color:#004C98; //#7CD2A6;
$success-notification:#009D66;
$error-notification:#BF3714;
$secondary-color:#95989A;
$terciary-color:#37495F;

$disabled:#EBEAE4;
$dark-gray:#6A6C6B;

$background-color:lighten(#EDEDED,3.5%);
$background-light:#FCFCFC;
$background-medium:#E4E4E4;
$background-dark:darken(#BFBDBD,4.5%);
$background-darker:darken(#BFBDBD,4.5%);
$highlight-1:$main-color;
$highlight-2:#983824;
$text-color:darken(#95989A, 15%);
$light-color:#FFFFFF;
$button-color: $highlight-1;
$chart-primary:#689BCD;
$chart-primary-darker:#00559F;
$chart-secondary:#CBCBCB;


.autocomplete {

    position: relative;
    display: inline-block;
    width: 100%;

    .dropdown-list {
        display: none;
        position: absolute;
        background-color: #f9f9f9;

        width: 100%;
        box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
        padding: 5px 5px;
        z-index: 1;
        .dropdown-option{
          width: auto;
          cursor:pointer;
          &:hover{
            background-color: darken($background-color,6%);
          }
        }
    }
    input{
      width:100%;
    }
    input:focus ~ .dropdown-list {
        display: block;
    }
    .dropdown-list:hover{
      display:block
    }
}
</style>

<script>
export default{
  props: ['value','options', 'attribute', 'beforeSearchMessage', 'searchingMessage', 'emptySearchResultsMessage', 'error'],
  data(){
    return {
      query: this.value,
      limit: 5,
    }
  },
  methods:{
      selectItem: function(item)
      {
          this.query = item[this.attribute];
          this.$emit('itemSelected', item);
      },
      setQuery: function(text) {
        this.query = text;
      }

  },
  created:function(){

  },
  mounted:function(){
  },
  updated: function(){
      this.$emit('updated', this.query);
  },
  watch:{

  },
  computed:{
    reducedOptions(){
      return this.limit ? this.options.slice(0,this.limit) : this.options
    }
  }
}
</script>