//
//  FontManager.swift
//  TriviaCast
//
//  Created by Brophy on 2/5/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

typealias TCFont = UIFont

extension TCFont {
    
    static let h1: TCFont = TCFont.systemFont(ofSize: h4.pointSize * 3.0)
    static let h2: TCFont = TCFont.systemFont(ofSize: h4.pointSize * 2.0)
    static let h3: TCFont = TCFont.systemFont(ofSize: h4.pointSize * 1.5)
    static let h4: TCFont = TCFont.systemFont(ofSize: 45.0)
    static let body: TCFont = TCFont.systemFont(ofSize: h4.pointSize)

}
