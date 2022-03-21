#!/bin/bash
#
# /*
#  * Copyright (C) 2022 The orange protocol Authors
#  * This file is part of The orange library.
#  *
#  * The Orange is free software: you can redistribute it and/or modify
#  * it under the terms of the GNU Lesser General Public License as published by
#  * the Free Software Foundation, either version 3 of the License, or
#  * (at your option) any later version.
#  *
#  * The orange is distributed in the hope that it will be useful,
#  * but WITHOUT ANY WARRANTY; without even the implied warranty of
#  * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#  * GNU Lesser General Public License for more details.
#  *
#  * You should have received a copy of the GNU Lesser General Public License
#  * along with The orange.  If not, see <http://www.gnu.org/licenses/>.
#  */
#

go get github.com/99designs/gqlgen/cmd@v0.14.0
go get github.com/99designs/gqlgen/internal/code@v0.14.0
go get github.com/99designs/gqlgen/internal/imports@v0.14.0
go run github.com/99designs/gqlgen generate