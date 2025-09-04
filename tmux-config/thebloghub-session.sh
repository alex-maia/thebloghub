#!/bin/bash

BASE_DIR=~/git/thebloghub/
SESSION=thebloghub

# Se a sessão não existir, cria
tmux has-session -t $SESSION 2>/dev/null
if [ $? != 0 ]; then
    # Cria sessão thebloghub com primeira janela 'thebloghub'
    tmux new-session -s $SESSION -n thebloghub -c "$BASE_DIR" -d

    # Cria segunda janela 'terminal'
    tmux new-window -t $SESSION:2 -n terminal -c "$BASE_DIR"

    # Cria segunda janela 'lazygit'
    tmux new-window -t $SESSION:3 -n lazygit -c "$BASE_DIR"

    # Envia o comando nvim para a primeira janela
    tmux send-keys -t $SESSION:thebloghub 'nvim' C-m
    
    # Envia o comando nvim para a terceira janela
    tmux send-keys -t $SESSION:lazygit 'lazygit' C-m
fi

# Attach automático
tmux attach-session -t $SESSION

# ( I use vim (BTW) 0_0 )

