unity-meta-checker
==================

This command detects missing correspondence between Unity asset files and .meta files.

## Installation
### Download binary

_TBW_

### Pull docker image

```sh
docker pull yuukiaria/unity-meta-checker
```

Or, just `docker run`

## Usage

### With release binary

```sh
/path/to/unity-meta-checker -p <UnityProjectPath>
```

### With docker

```sh
docker run --rm -v <UnityProjectPath>:/workspace:ro,cached yuukiaria/unity-meta-checker
```

## Commad Line Options

### -h, --help

Display help and exit.

### -p _path_, --project=_path_

Specify a path to Unity project, containing `Assets` folder.

### -e, --raise-error

Exit with non zero code when any paths are reported.

If this option is not specified, exit with zero even if some paths are reported.

### -o _path_, --output=_path_

Specify a path to file that reports be written to.

### -t _path_, --template=_path_

Specify a path to template file used to render reports.
Templates are processed by [Golang's text/template package](https://golang.org/pkg/text/template/).

For example, use _custom.yml.tpl_ with the following contents:

```
danglingMetaFiles:
{{- range .DanglingMetaPaths }}
  - {{ . }}
{{- else }} [] {{- end }}
metalessAssetPaths:
{{- range .MetalessAssetPaths }}
  - {{ . }}
{{- else }} [] {{- end }}
```

Then reports will be rendered like as:

```yml
danglingMetaFiles:
  - abc/def/ghi.meta
metalessAssetPaths: []
```
