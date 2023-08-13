import React from "react";
import theme from "styles/extendTheme";

import { Box, Link } from "@chakra-ui/react";
import { NavigateFunction } from "react-router-dom";

import { Account } from "types/Account";
import { Create as CreateQuery } from "postAPI/accounts/Create";
import { Login as LoginQuery } from "postAPI/accounts/Login";

import Input from "components/Input";
import RoundButton from "components/RoundButton";

import styles from "./SignupPage.module.scss";
import { RegistrationCard } from "types/RegistrationCard";

type SignUpProps = {
    navigate: NavigateFunction
}


class SignUpPage extends React.Component<SignUpProps> {
    registrationCard: RegistrationCard = {
            first_name: "",
            last_name: "",
            email: "",
            username: "",
            password: ""
    }

    repPassword: string = ""

    setfirst_name(val: string) {
        this.registrationCard.first_name = val
    }
    setlast_name(val: string) {
        this.registrationCard.last_name = val
    }
    setEmail(val: string) {
        this.registrationCard.email = val
    }
    setLogin(val: string) {
        this.registrationCard.username = val
    }
    setPassword(val: string) {
        this.registrationCard.password = val
    }
    setRepPassword(val: string) {
        this.repPassword = val
    }

    highlightNotMatch() {
        let node1 = document.getElementsByName("password")[0]
        let node2 = document.getElementsByName("rep-password")[0]

        if (node1.parentElement && node2.parentElement) {
            node1.parentElement.style.borderColor = theme.colors["title"]
            node2.parentElement.style.borderColor = theme.colors["title"]
        }

        var title = document.getElementById("undertitle")
        if (title)
            title.innerText = "Пароли не совпадают!"
    }

    async submit(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        if (this.registrationCard.password !== this.repPassword)
            return this.highlightNotMatch()

        //e.currentTarget.disabled = true
        var data = await CreateQuery(this.registrationCard)
        if (data.status === 200) {
            let acc: Account = {
                username: this.registrationCard.username,
                password: this.registrationCard.password
            }

            await LoginQuery(acc);
            window.location.href = '/';
        } else {
            //e.currentTarget.disabled = false
            var title = document.getElementById("undertitle")
            if (title)
                title.innerText = "Ошибка создания аккаунта!"
        };
    }

    render() {
        return <Box className={styles.login_page}>
            <Box className={styles.input_div}>
                <Input name="first_name" placeholder="Введите имя"
                onInput={event => this.setfirst_name(event.currentTarget.value)}/>
                <Input name="last_name" placeholder="Введите фамилию"
                onInput={event => this.setlast_name(event.currentTarget.value)}/>
                <Input name="email" placeholder="Введите электронную почту"
                onInput={event => this.setEmail(event.currentTarget.value)}/>
                <Input name="username" placeholder="Введите логин"
                onInput={event => this.setLogin(event.currentTarget.value)}/>
                {/*<Input name="login" placeholder="Введите телефон" */}
                {/*onInput={event => this.setMobilePhone(event.currentTarget.value)}/>*/}
                <Input name="password" type="password" placeholder="Введите пароль"
                onInput={event => this.setPassword(event.currentTarget.value)}/>
                <Input name="rep-password" type="password" placeholder="Повторите пароль"
                onInput={event => this.setRepPassword(event.currentTarget.value)}/>
            </Box>

            <Box className={styles.input_div}>
                <RoundButton type="submit" onClick={event => this.submit(event)}>
                    Создать аккаунт
                </RoundButton>
                <Link href="/authorize">Войти</Link>
            </Box>
        </Box>
    }
}

export default SignUpPage;