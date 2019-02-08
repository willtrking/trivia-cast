//
//  Layout.swift
//  TriviaCast
//
//  Created by Brophy on 2/5/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit
import SnapKit

typealias SpacingUnit = CGFloat


extension SpacingUnit {
    static let u0 = SpacingUnit(0.0)
    static let u1 = SpacingUnit(10.0)
    static let u2 = SpacingUnit(u1 * 2.0)
    static let u3 = SpacingUnit(u1 * 3.0)
    static let u4 = SpacingUnit(u1 * 4.0)
    static let u5 = SpacingUnit(u1 * 5.0)
}

//extension CGFloat {
//    /// Creates a new instance from the given value, rounded to the closest
//    /// possible representation.
//    ///
//    /// - Parameter value: A SpacingUnit value to be converted.
//    init(_ value: SpacingUnit) {
//        self.init(value.rawValue)
//    }
//}
//
//extension SpacingUnit: ConstraintOffsetTarget {}
