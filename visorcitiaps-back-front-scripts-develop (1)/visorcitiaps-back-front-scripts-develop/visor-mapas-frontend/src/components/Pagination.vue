<template>

  <div class="pagination">
    <li class="page-item">
      <button
      class="page-link"
      type="button"
      @click="onClickFirstPage"
      :disabled="isInFirstPage"
      >
      <font-awesome-icon icon="angle-double-left"></font-awesome-icon>
      </button>
    </li>

    <li class="page-item">
      <button
      class="page-link"
      type="button"
      @click="onClickPreviousPage"
      :disabled="isInFirstPage"
      >
      <font-awesome-icon icon="angle-left"></font-awesome-icon>
      </button>
    </li>

    <li
      v-for="page in pages"
      v-bind:key="page.index"
      class="page-item"
      :class="{ activepage: isPageActive(page.name) }"
    >
      <button
      class="page-link"
      type="button"
      @click="onClickPage(page.name)"
      :disabled="page.isDisabled"
      :class="{ activepage: isPageActive(page.name) }"
      >
      {{ page.name }}
      </button>
    </li>

    <li class="page-item">
      <button
      class="page-link"
      type="button"
      @click="onClickNextPage"
      :disabled="isInLastPage"
      >
      <font-awesome-icon icon="angle-right"></font-awesome-icon>
      </button>
    </li>

    <li class="page-item">
      <button
      class="page-link"
      type="button"
      @click="onClickLastPage"
      :disabled="isInLastPage"
      >
      <font-awesome-icon icon="angle-double-right"></font-awesome-icon>
      </button>
    </li>
  </div>

</template>


<style lang="scss" scoped>
.pagination{
  display: flex;
  list-style-type: none;
  justify-content: center;
}
.page-item{
  border: 2px solid rgba(0, 0, 0, 0.25);
  .page-link{
    border: none;
    background: white;
    font-size: 1em;
    font-weight: bold;
    padding: 0.4em 0.8em;
  }
}
.page-item:hover{
  background: #f7e200;
  .page-link{
    background: #f7e200;
  }
}
.page-item:first-child{
  border-top-left-radius: 6px;
  border-bottom-left-radius: 6px;
}
.page-item:last-child{
  border-top-right-radius: 6px;
  border-bottom-right-radius: 6px;
}
.activepage {
  background-color: #03440c;
  color: white;
  .page-link{
    background: #03440c;
  }
}

@media only screen and (max-width: 640px) {
  .page-item{
    .page-link{
      font-size: 1em;
      padding: 0.25em 0.5em;
    }
  }
}
</style>



<script>
export default {
  props: {
    totalPages: {
      type: Number,
      required: true
    },
    currentPage: {
      type: Number,
      required: true
    },
    pageSize: {
      type: Number,
      required: true
    }
  },
  computed: {
    maxVisibleButtons: function() {
        if(parseInt(this.totalPages, 10) < 3)
        {
            return this.totalPages;
        }
        return 3;
    },
    startPage() {
      if (this.currentPage === 1) {
        return 1;
      }
      if (this.currentPage === this.totalPages) {
        return this.totalPages - this.maxVisibleButtons + 1;
      }
      return this.currentPage - 1;
    },
    endPage() {
      return Math.min(this.startPage + this.maxVisibleButtons - 1, this.totalPages);
    },
    pages() {
      const range = [];
      for (let i = this.startPage; i <= this.endPage; i+= 1 ) {
        range.push({
          name: i,
          isDisabled: i === this.currentPage
        });
      }
      return range;
    },
    isInFirstPage() {
      return this.currentPage === 1;
    },
    isInLastPage() {
      return this.currentPage === this.totalPages;
    }
  },
  methods: {
    onClickFirstPage() {
      this.$emit('pagechanged', 1);
    },
    onClickPreviousPage() {
      this.$emit('pagechanged', this.currentPage - 1);
    },
    onClickPage(page) {
      this.$emit('pagechanged', page);
    },
    onClickNextPage() {
      this.$emit('pagechanged', this.currentPage + 1);
    },
    onClickLastPage() {
      this.$emit('pagechanged', this.totalPages);
    },
    isPageActive(page) {
      return this.currentPage === page;
    }
  }
};
</script>
