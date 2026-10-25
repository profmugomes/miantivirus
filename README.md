# MiAntivirus

MiAntivirus é uma interface gráfica para o ClamAV que permite verificar seu computador em busca de vírus e atualizar facilmente o banco de dados de definições de vírus por meio de uma interface amigável.

## Capturas de tela

<img width="660" height="506" alt="image" src="https://github.com/user-attachments/assets/ea248e4d-b97a-4ca5-88b3-05077318fb64" />

<img width="564" height="674" alt="image" src="https://github.com/user-attachments/assets/be60d4d3-91f9-4eaf-8e97-c4188e31e27b" />

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

## Information

 - [Repository MiAntivirus](https://github.com/profmugomes/miantivirus)

## System Requirement

### Linux

- Ubuntu 26.04 or higher

## 👤 Autor

**Murilo Gomes**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

The MiAntivirus is provided under:

[SPDX-License-Identifier: GPL-2.0-only](https://spdx.org/licenses/GPL-2.0-only.html)

Beign under the terms of the GNU General Public License version 2 only.

All contributions to the MiAntivirus are subject to this license.
