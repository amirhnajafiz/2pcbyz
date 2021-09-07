import Home from './components/Home';
import Categories from './screens/category/Categories';

export default {
    mode: 'history',
    routes: [
        {
            path: '/',
            component: Home
        },
        {
            path: '/admin/categories',
            component: Categories
        }
    ]
}