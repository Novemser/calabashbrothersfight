import Vue from 'vue'
import Router from 'vue-router'
import IndexPage from '@/components/Catalog'
import Level from '@/components/Level'

Vue.use(Router)

export default new Router({
	routes: [
		{
			path: '/',
			name: 'Index',
			component: IndexPage
		},
		{
			path: '/level/:level',
			name: 'level',
			component: Level
		}
	]
})
