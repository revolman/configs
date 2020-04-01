
# Tmux + zsh
## Настройка терминала tmux и шела zsh

```apt install -y zsh```

### oh-my-zsh
Темы для zsh

https://github.com/ohmyzsh/ohmyzsh
Я редактировал эту:
<home>/.oh-my-zsh/themes/robbyrussell.zsh-theme

```
function prompt_char {
        if [ $UID -eq 0 ]; then echo " #"; fi
}

local ret_status="%(?:%{$fg_bold[green]%}➤➤$(prompt_char) :%{$fg_bold[red]%}➤➤$(prompt_char) )"
PROMPT='%{$fg_bold[white]%}$(whoami):%{$reset_color%} %{$fg_bold[cyan]%}%c%{$reset_color%} $(git_prompt_info) ${ret_status}%{$reset_color%}'
#PROMPT+=' %{$fg[cyan]%}%c%{$reset_color%} $(git_prompt_info)'

ZSH_THEME_GIT_PROMPT_PREFIX="%{$fg_bold[blue]%}git:(%{$fg[red]%}"
ZSH_THEME_GIT_PROMPT_SUFFIX="%{$reset_color%} "
ZSH_THEME_GIT_PROMPT_DIRTY="%{$fg[blue]%}) %{$fg[yellow]%}✱"
ZSH_THEME_GIT_PROMPT_CLEAN="%{$fg[blue]%})"
```
Что бы применить изменения:
```source .zshrc```

### zsh-autosuggestions
Автодополнение комманд

https://github.com/zsh-users/zsh-autosuggestions/blob/master/INSTALL.md

После установки включить плагин в .zshrc:
```plugins=(git kubectl zsh-autosuggestions)```
git и kubelet - другие плагины.

### Tmux
Установка и настройка tmux

```apt install -y tmux```

Дале в терминале tmux:
```tmux show -g | cat > /home/<username>/.tmux.conf```

Добавить в .tmux.conf:
```
source /usr/share/powerline/bindings/tmux/powerline.conf
set -g mouse on
```
Исправить цветовую схему zsh:
```echo "export TERM=xterm-256color" >> .zshrc```

### Powerline
Панель состояния для Tmux

```apt install -y powerline```

Установить шрифты:
```
wget https://github.com/powerline/powerline/raw/develop/font/PowerlineSymbols.otf
wget https://github.com/powerline/powerline/raw/develop/font/10-powerline-symbols.conf
mv PowerlineSymbols.otf /usr/share/fonts/
fc-cache -vf /usr/share/fonts/
mv 10-powerline-symbols.conf /etc/fonts/conf.d/
```

Source: 
```/usr/share/powerline/bindings/tmux/powerline.conf```
