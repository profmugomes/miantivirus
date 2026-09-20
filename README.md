# MiAntivirus

[![License](https://img.shields.io/badge/license-PolyForm%20Perimeter%201.0.1-5351FB)](LICENSE.md)

MiAntivirus é uma interface gráfica para o ClamAV que permite verificar seu computador em busca de vírus e atualizar facilmente o banco de dados de definições de vírus por meio de uma interface amigável.

## Capturas de tela

<img height="300" alt="image" src="https://github.com/user-attachments/assets/7b2299fa-7515-4be8-9b00-604ed03b73da" />

<img height="300" alt="image" src="https://github.com/user-attachments/assets/9e362326-73e4-4441-8b90-15bfd789fce7" />

## Contribuição

1. Relate bugs ou envie sugestões: https://github.com/profmugomes/miantivirus/issues

2. Envie traduções para outros idiomas: baixe o arquivo `.po`, traduza-o e envie-o por meio de um Pull Request.

3. Se quiser, dê uma "estrela" ao MiAntivirus.

4. Apoie financeiramente o projeto por meio do GitHub Sponsors ou de outras formas de apoio.

## Apoio

* https://github.com/sponsors/profmugomes


## Instalação

### DEB

Clique duas vezes no pacote deb e clique em instalar, ou execute o comando no terminal:

```bash
sudo dpkg -i miantivirus*.deb

sudo apt install -f
```

### Integridade

Para verificar a integridade, copie o hash fornecido ao lado da versão baixada e use o [MiCheckHash](https://github.com/profmugomes/micheckhash/releases) (interface gráfica) para verificar sua integridade, ou use o terminal.

Exemplo:

```bash
echo "12f95d1ba9b46f5713d8010963c4c782e315b7985027c44c6e292ede69454301 miantivirus_4.0.0_all.deb" | sha256sum -c
```

Se "Success" for exibido, o arquivo foi baixado corretamente.

## Uso

O MiAntivirus permite verificar vários arquivos e pastas; basta clicar em "Adicionar" e escolher quais deseja adicionar. Você pode adicionar ambos os tipos.

Antes da verificação, sempre recomendo atualizar o banco de dados do ClamAV. Para isso, clique em "Atualizar banco de dados" no menu Ferramentas. Quando a atualização for concluída, uma mensagem "Concluído" será exibida.

Em Opções, no menu Ferramentas, você pode ativar e desativar os recursos de verificação e adicionar e remover pastas e arquivos que deseja ignorar durante a verificação.

Ao clicar em "Verificar", o ClamAV será carregado. Esse processo pode levar algum tempo (essa lentidão no carregamento do banco de dados ocorre devido ao próprio ClamAV). Após o carregamento, as pastas e os arquivos selecionados serão verificados. Quando a análise for concluída, você poderá remover arquivos ou pastas infectados, caso algum vírus seja detectado.

**Atenção:** o ClamAV pode apresentar alguns falsos positivos, portanto, considere se realmente deseja excluí-los. Essa exclusão é permanente e não pode ser recuperada. Portanto, tenha cuidado ao excluir um arquivo ou pasta.

Para verificar se há novas atualizações do MiAntivirus, clique em Verificar atualizações no menu Sobre.

### Comandos Externos

O MiAntivirus utiliza o ClamAV e o pkexec por meio de seus comandos disponibilizados pelo sistema operacional.

Esses recursos são executados externamente pelo MiAntivirus e não fazem parte do código ou dos binários distribuídos pelo projeto. Dessa forma, o funcionamento desses componentes depende da instalação e configuração adequada no sistema do usuário.

## Limitações/Bugs

O software pode conter algumas limitações ou bugs, portanto, é muito importante utilizar os canais oficiais de contato para relatar bugs.

## 👤 Autor

**Murilo Gomes**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes <profmugomes.com.br>. All Rights Reserved.

This project is licensed under the PolyForm Perimeter License 1.0.1.

### Summary

This software is available for commercial and noncommercial use, subject to the terms of the PolyForm Perimeter License 1.0.1.

You may:

* ✔ Use the software for commercial and noncommercial purposes.
* ✔ Inspect and study the source code.
* ✔ Modify the software.
* ✔ Create derivative works based on the software.
* ✔ Redistribute the software and permitted modifications.

You may not:

* ✖ Provide a product that competes with the software.

See the full license terms at LICENSE.md.

This summary is provided for convenience only and does not replace or modify the full license terms.
