import Vue from 'vue'
import Router from 'vue-router'
import IndexPage from '@/components/Catalog'
import Beginning from '@/components/Section/Beginning'

Vue.use(Router)

export default new Router({
	routes: [
		{
			path: '/',
			name: 'Index',
			component: IndexPage
		},
		{
			path: '/section/beginning',
			name: 'Beginning',
			component: Beginning
		}
	]
})
