<script setup lang="ts">
import { ref, onMounted } from 'vue'

const keyword = ref<string>('')
const minPrice = ref<number>()
const maxPrice = ref<number>()
const sort = ref<string>('default')
const error = ref<string>()

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
    affiliateUrl: string;
    itemName: string;
    itemPrice: number;
    itemCode: string;
    shopName: string;
    shopAffiliateUrl: string;
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

const rakutenSearchResults = ref<RakutenItemDetail[]>([])
const amazonSearchResults = ref<AmazonItemDetail[]>([])

interface SearchRequest {
  keyword: string;
  minPrice?: number;
  maxPrice?: number;
  sort: string;
}

function amazonSearch(request: SearchRequest) {
  fetch('/api/amazon/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(request)
  })
    .then((resp) => resp.json())
    .then((data) => {
      amazonSearchResults.value = data.SearchResult.Items
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

function rakutenSearch(SearchRequest: SearchRequest) {
  fetch('/api/rakuten/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(SearchRequest)
  })
    .then((resp) => resp.json())
    .then((data) => {
      rakutenSearchResults.value = data.Items
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

function search() {
  if (!keyword.value) {
    error.value = '商品名を入力してください'
    return
  }
  if (minPrice.value && maxPrice.value && minPrice.value > maxPrice.value) {
    error.value = '最低価格が最高価格より高いです'
    return
  }
  if (minPrice.value && minPrice.value < 0) {
    error.value = '最低価格が0より小さいです'
    return
  }
  if (maxPrice.value && maxPrice.value < 0) {
    error.value = '最高価格が0より小さいです'
    return
  }
  error.value = ''
  const request: SearchRequest = {
    keyword: keyword.value,
    minPrice: minPrice.value,
    maxPrice: maxPrice.value,
    sort: sort.value
  }
  amazonSearch(request)
  rakutenSearch(request)
}

</script>

<template>
  <div class="container">
    <h1>一括商品 商品検索</h1>
    <div class="search-form">
      <input v-model="keyword" placeholder="商品名を入力" class="search-input" />
      <input v-model="minPrice" type="number" placeholder="最低価格" class="price-input" />
      <input v-model="maxPrice" type="number" placeholder="最高価格" class="price-input" />
      <select v-model="sort" class="sort-select">
        <option value="default">標準</option>
        <option value="price_asc">価格順(安い順)</option>
        <option value="price_desc">価格順(高い順)</option>
      </select>
      <button @click="search" class="search-button">検索</button><br>
    </div>
    <div class="error" v-if="error">
      <p class="error-message">{{ error }}</p>
    </div>

    <div class="results-container">
      <div class="result-section">
        <h2>楽天市場</h2>
        <ul class="result-list">
          <li v-for="(item, index) in rakutenSearchResults" :key="index" class="item">
            <a :href="item.Item.affiliateUrl" class="item-link">
              <img :src="item.Item.mediumImageUrls[0]?.imageUrl" class="item-image">
              <div class="item-details">
                <h3>{{ item.Item.itemName }}</h3>
                <p>{{ item.Item.itemPrice }}円</p>
                <p>ショップ: <a :href="item.Item.shopAffiliateUrl" class="shop-link">{{ item.Item.shopName }}</a></p>
              </div>
            </a>
          </li>
        </ul>
      </div>
      <div class="result-section">
        <h2>Amazon</h2>
        <ul class="result-list">
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
  </div>
</template>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}

.search-form {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.search-input,
.price-input,
.sort-select,
.search-button {
  flex: 1;
  padding: 10px 15px;
  border: 2px solid #ccc;
  border-radius: 6px;
  outline: none;
  transition: border 0.3s ease;
}

.search-input:focus,
.price-input:focus,
.sort-select:focus {
  border-color: #0073e6;
}

.search-button {
  background-color: #0073e6;
  color: #fff;
  cursor: pointer;
  border: none;
  transition: background 0.3s ease;
}

.search-button:hover {
  background-color: #005bb5;
}

.error {
  margin-bottom: 20px;
}

.error-message {
  color: #ff0000;
}

.results-container {
  display: flex;
  gap: 20px;
}

.result-section {
  flex: 1;
  background-color: #fff;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.result-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.item {
  border-bottom: 1px solid #eee;
  padding: 12px;
  display: flex;
  align-items: center;
}

.item:last-child {
  border-bottom: none;
}

.item-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: inherit;
}

.item-image {
  width: 100px;
  border-radius: 6px;
  margin-right: 15px;
}

.item-details h3 {
  margin: 0 0 5px 0;
  font-size: 16px;
  color: #333;
}

.item-details p {
  margin: 0;
  font-size: 14px;
  color: #666;
}

.shop-link {
  color: #0073e6;
  text-decoration: underline;
}
</style>
