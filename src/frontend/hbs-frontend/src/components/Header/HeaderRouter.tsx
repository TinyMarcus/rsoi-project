import React from "react";
import { Routes, Route } from "react-router";
import { BrowserRouter } from "react-router-dom";
import Header from ".";
import CategoryHeader from "./Category";
import RecipeHeader from "./Recipe";
import SearchHeader from "./Search";
import UserHeader from "./User";


export const HeaderRouter: React.FC<{}> = () => {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<Header title="Все отели"/>}/>
            <Route path="/users" element={<SearchHeader title="Все пользователи"/>}/>
            <Route path="/me/reservations" element={<Header subtitle="Бронирования" title="Мои" />}/>
            <Route path="/statistics" element={<Header subtitle="Общая" title="Статистика" />}/>

            <Route path="/accounts/:login/recipes" element={<UserHeader subtitle="Автор" title="" />}/>
            <Route path="/accounts/:login/likes" element={<UserHeader subtitle="Понравилось" title="" />}/>

            <Route path="/authorize" element={<Header title="Вход" undertitle="Добро пожаловать. Снова." />}/>
            <Route path="/register" element={<Header title="Регистрация" undertitle="Чтобы получить доступ к бронированиям!" />}/>
            
            <Route path="/recipes/:id" element={<RecipeHeader title=""/>}/>
            <Route path="/categories/:title" element={<CategoryHeader subtitle="Категория" title=""/>}/>

            <Route path="*" element={<Header title="Страница не найдена"/>}/>
        </Routes>
    </BrowserRouter>
}
