<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

const keyword = ref<string>('')
const minPrice = ref<number>()
const maxPrice = ref<number>()
const sort = ref<string>('default')
const error = ref<string>()

// ページング用の状態
const currentPage = ref<number>(1)
const pageSize = ref<number>(10)
const totalItems = ref<number>(0)
const totalPages = ref<number>(0)

// 表示するページ番号を計算するcomputedプロパティ
const displayedPages = computed(() => {
  const pages: number[] = []
  const maxVisiblePages = 5
  let startPage = Math.max(1, currentPage.value - Math.floor(maxVisiblePages / 2))
  let endPage = startPage + maxVisiblePages - 1
  
  if (endPage > totalPages.value) {
    endPage = totalPages.value
    startPage = Math.max(1, endPage - maxVisiblePages + 1)
  }
  
  for (let i = startPage; i <= endPage; i++) {
    pages.push(i)
  }
  
  return pages
})

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
  page?: number;
  size?: number;
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
      // ページング情報の更新
      if (data.SearchResult) {
        // Amazonの場合は個別のカウントを使用
        amazonPaging.value.totalItems = data.SearchResult.TotalResultCount || 0
        amazonPaging.value.totalPages = Math.ceil(amazonPaging.value.totalItems / pageSize.value)
        amazonSearchResults.value = data.SearchResult.Items || []
      }
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

// 各サービスのページング情報を別々に管理
const rakutenPaging = ref({
  totalItems: 0,
  totalPages: 0
})

const amazonPaging = ref({
  totalItems: 0,
  totalPages: 0
})

function rakutenSearch(request: SearchRequest) {
  fetch('/api/rakuten/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(request)
  })
    .then((resp) => resp.json())
    .then((data) => {
      // ページング情報の更新
      if (data) {
        rakutenPaging.value.totalItems = data.count || 0
        rakutenPaging.value.totalPages = Math.ceil(rakutenPaging.value.totalItems / pageSize.value)
        rakutenSearchResults.value = data.Items || []
      }
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}

function search(page: number = 1) {
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
    sort: sort.value,
    page: page,
    size: pageSize.value
  }
  currentPage.value = page
  amazonSearch(request)
  rakutenSearch(request)
}

function handleSearchClick() {
  search()
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
      <button @click="handleSearchClick" class="search-button">検索</button><br>
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
    
    <!-- ページングコントロール -->
    <div class="pagination-section">
      <div class="pagination-controls" v-if="rakutenPaging.totalPages > 1">
        <h3>楽天市場 ページング</h3>
        <button 
          @click="search(1)" 
          :disabled="currentPage === 1"
          class="pagination-button"
        >
          最初へ
        </button>
        
        <button 
          @click="search(currentPage - 1)" 
          :disabled="currentPage === 1"
          class="pagination-button"
        >
          前へ
        </button>
        
        <!-- ページ番号の表示（現在のページを中心に表示） -->
        <span 
          v-for="page in displayedPages" 
          :key="page"
          @click="search(page)"
          :class="{ 'current-page': page === currentPage }"
          class="page-number"
        >
          {{ page }}
        </span>
        
        <button 
          @click="search(currentPage + 1)" 
          :disabled="currentPage === rakutenPaging.totalPages"
          class="pagination-button"
        >
          次へ
        </button>
        
        <button 
          @click="search(rakutenPaging.totalPages)" 
          :disabled="currentPage === rakutenPaging.totalPages"
          class="pagination-button"
        >
          最後へ
        </button>
        <p>全{{ rakutenPaging.totalItems }}件中 {{ currentPage * pageSize - pageSize + 1 }}-{{ Math.min(currentPage * pageSize, rakutenPaging.totalItems) }}件を表示</p>
      </div>
      
      <div class="pagination-controls" v-if="amazonPaging.totalPages > 1">
        <h3>Amazon ページング</h3>
        <button 
          @click="search(1)" 
          :disabled="currentPage === 1"
          class="pagination-button"
        >
          最初へ
        </button>
        
        <button 
          @click="search(currentPage - 1)" 
          :disabled="currentPage === 1"
          class="pagination-button"
        >
          前へ
        </button>
        
        <!-- ページ番号の表示（現在のページを中心に表示） -->
        <span 
          v-for="page in displayedPages" 
          :key="page"
          @click="search(page)"
          :class="{ 'current-page': page === currentPage }"
          class="page-number"
        >
          {{ page }}
        </span>
        
        <button 
          @click="search(currentPage + 1)" 
          :disabled="currentPage === amazonPaging.totalPages"
          class="pagination-button"
        >
          次へ
        </button>
        
        <button 
          @click="search(amazonPaging.totalPages)" 
          :disabled="currentPage === amazonPaging.totalPages"
          class="pagination-button"
        >
          最後へ
        </button>
        <p>全{{ amazonPaging.totalItems }}件中 {{ currentPage * pageSize - pageSize + 1 }}-{{ Math.min(currentPage * pageSize, amazonPaging.totalItems) }}件を表示</p>
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

/* ページングコントロールのスタイル */
.pagination-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-top: 30px;
  padding: 20px 0;
}

.pagination-button {
  padding: 8px 16px;
  background-color: #0073e6;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.pagination-button:hover:not(:disabled) {
  background-color: #005bb5;
}

.pagination-button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.page-number {
  padding: 8px 12px;
  cursor: pointer;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.page-number:hover {
  background-color: #e0e0e0;
}

.current-page {
  background-color: #0073e6;
  color: white;
}
</style>
