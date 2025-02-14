<script setup lang="ts">
import { ref, onMounted } from 'vue'

const keyword = ref<string>('')

onMounted(() => {
  const params = new URLSearchParams(window.location.search)
  const initialKeyword = params.get('keywords')
  if (initialKeyword) {
    keyword.value = initialKeyword
    search()
  }
})

interface RakutenItemDetail {
  Item: {
    mediumImageUrls: { imageUrl: string }[];
    itemUrl: string;
    itemName: string;
    itemPrice: number;
    itemCode: string;
  }
}

interface AmazonItemDetail {
  ASIN: string;
  DetailPageURL: string;
  Images: {
    Primary: {
      Medium: { URL: string };
    };
  };
  ItemInfo: {
    Title: {
      DisplayValue: string;
    };
    ByLineInfo: {
      Brand: {
        DisplayValue: string;
      };
    };
    Classifications: {
      ProductGroup: {
        DisplayValue: string;
      };
    };
    Features: {
      DisplayValues: string[];
    };
  };
  Offers: {
    Listings: {
      Price: {
        DisplayAmount: string;
        Price: number;
      };
    }[];
  };
}

//const keyword = ref('')
const rakutenSearchResults = ref<RakutenItemDetail[]>([])
const amazonSearchResults = ref<AmazonItemDetail[]>([])

function amazonSearch() {
  const formData = new FormData()
  formData.append('keyword', keyword.value)

  fetch('/api/amazon/search', {
    method: 'POST',
    body: formData
  })
    .then((resp) => resp.json())
    .then((data) => {
      amazonSearchResults.value = data.SearchResult.Items
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

function rakutenSearch() {
  const formData = new FormData()
  formData.append('keyword', keyword.value)

  fetch('/api/rakuten/search', {
    method: 'POST',
    body: formData
  })
    .then((resp) => resp.json())
    .then((data) => {
      rakutenSearchResults.value = [...data.Items]
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

function search() {
  amazonSearch()
  rakutenSearch()
}

</script>

<template>
  <div>
    <h1>一括商品 商品検索</h1>
    <input v-model="keyword" placeholder="商品名を入力" class="search-input" />
    <button @click="search" class="search-button">検索</button>
    <div class="results-container">
      <ul class="result-list rakuten-results">
        <h2>楽天市場</h2>
        <li v-for="(item, index) in rakutenSearchResults" :key="index" class="item">
          <a :href="item.Item.itemUrl" class="item-link">
            <img :src="item.Item.mediumImageUrls[0]?.imageUrl" class="item-image">
            <div class="item-details">
              <h3>{{ item.Item.itemName }}</h3>
              <p>{{ item.Item.itemPrice }}円</p>
            </div>
          </a>
        </li>
      </ul>
      <ul class="result-list amazon-results">
        <h2>Amazon</h2>
        <li v-for="(item, index) in amazonSearchResults" :key="index" class="item">
          <a :href="item.DetailPageURL" class="item-link">
            <img :src="item.Images.Primary.Medium.URL" class="item-image">
            <div class="item-details">
              <h3>{{ item.ItemInfo.Title.DisplayValue }}</h3>
              <p>{{ item.Offers?.Listings[0]?.Price.DisplayAmount }}</p>
            </div>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.search-input,
.search-button {
  padding: 8px 10px;
  margin: 10px 5px;
  border: 2px solid #ccc;
  border-radius: 4px;
}

.search-button {
  cursor: pointer;
  background-color: #f0f0f0;
}

.results-container {
  display: flex;
  justify-content: space-between;
}

.result-list {
  flex: 1;
  list-style: none;
  padding: 0;
  margin: 0 10px;
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
