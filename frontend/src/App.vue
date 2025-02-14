<script setup lang="ts">
import { ref } from 'vue'

interface ItemDetail {
  Item: {
    mediumImageUrls: { imageUrl: string }[];
    itemUrl: string;
    itemName: string;
    itemPrice: number;
    itemCode: string;
  }
}

const keyword = ref('')
const searchResults = ref<ItemDetail[]>([])

function search() {
  const formData = new FormData()
  formData.append('keyword', keyword.value)

  fetch('/api/search', {
    method: 'POST',
    body: formData
  })
    .then((resp) => resp.json())
    .then((data) => {
      searchResults.value = [...data.Items]
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

</script>

<template>
  <div>
    <input v-model="keyword" placeholder="商品名を入力" class="search-input" />
    <button @click="search" class="search-button">検索</button>
    <ul class="result-list">
      <li v-for="(item, index) in searchResults" :key="index" class="item">
        <a :href="item.Item.itemUrl" class="item-link">
          <img :src="item.Item.mediumImageUrls[0].imageUrl" class="item-image">
          <div class="item-details">
            <h3>{{ item.Item.itemName }}</h3>
            <p>{{ item.Item.itemPrice }}円</p>
          </div>
        </a>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.search-input, .search-button {
  padding: 8px 10px;
  margin: 10px 5px;
  border: 2px solid #ccc;
  border-radius: 4px;
}

.search-button {
  cursor: pointer;
  background-color: #f0f0f0;
}

.result-list {
  list-style: none;
  padding: 0;
}

.item {
  border-bottom: 1px solid #eee;
  padding: 10px;
  display: flex;
  align-items: center;
}

.item-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: inherit;
}

.item-image {
  width: 100px;
  margin-right: 20px;
}

.item-details h3 {
  margin: 0 0 10px 0;
  font-size: 16px;
}

.item-details p {
  margin: 0;
  font-size: 14px;
}
</style>
