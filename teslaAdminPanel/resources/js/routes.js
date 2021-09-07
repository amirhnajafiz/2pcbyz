import Home from './components/Home';
import Categories from './screens/category/Categories';
import AddCategory from './screens/category/AddCategory';

export default {
    mode: 'history',
    routes: [
        {
            path: '/admin',
            component: Home
        },
        {
            path: '/admin/categories',
            component: Categories
        },
        {
            path: '/admin/addCategory',
            component: AddCategory
        }
    ]
}